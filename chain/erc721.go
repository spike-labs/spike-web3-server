package chain

import (
	"container/list"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"spike-frame/cache"
	chain "spike-frame/chain/abi"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/model"
	"spike-frame/util"
	"strings"
	"sync"
)

func (e *ERC721Listener) Notify(event cache.ClearEvent) {
	e.observerLk.Lock()
	defer e.observerLk.Unlock()

	for o := e.observers.Front(); o != nil; o = o.Next() {
		if ov, ok := o.Value.(cache.Observer); ok {
			ov.Update(event)
		}
	}
}

func (e *ERC721Listener) AttachObserver(observer cache.Observer) {
	e.observerLk.Lock()
	defer e.observerLk.Unlock()

	o := e.observers.Front()
	if o == nil {
		e.observers.PushBack(observer)
	}
}

func (e *ERC721Listener) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(constant.EmptyAddress) == strings.ToLower(fromAddr) {
		return true, constant.GAMENFT_TRANSFER
	}

	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(toAddr) {
		return true, constant.GAMENFT_IMPORT
	}
	return true, constant.GAMENFT_TRANSFER
}

type ERC721Listener struct {
	contractAddr   string
	tokenType      model.TokenType
	newBlockNotify util.DataChannel
	ec             *ethclient.Client
	abi            abi.ABI
	errorHandler   chan ErrMsg
	observers      *list.List
	observerLk     sync.Mutex
}

func newERC721Listener(contractAddr string, tokenType model.TokenType, ec *ethclient.Client, newBlockNotify util.DataChannel, abi abi.ABI, errorHandler chan ErrMsg) *ERC721Listener {
	e := &ERC721Listener{
		contractAddr:   contractAddr,
		tokenType:      tokenType,
		newBlockNotify: newBlockNotify,
		ec:             ec,
		abi:            abi,
		errorHandler:   errorHandler,
		observers:      list.New(),
	}
	e.AttachObserver(cache.NewManager(constant.RedisClient))
	return e
}

func (el *ERC721Listener) run() {
	go el.NewEventFilter()
}

func (el *ERC721Listener) NewEventFilter() error {
	for {
		select {
		case de := <-el.newBlockNotify:
			height := de.Data.(*big.Int)
			el.handlePastBlock(height, height)
		}
	}
}

func (el *ERC721Listener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("nft past event filter, fromBlock : %d, toBlock : %d ", fromBlockNum, toBlockNum)
	ethClient := el.ec
	contractAddress := common.HexToAddress(el.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		el.errorHandler <- ErrMsg{
			tp:   el.tokenType,
			from: fromBlockNum,
			to:   toBlockNum,
		}
		log.Errorf("nft subscribe event log, from: %d,to: %d,err : %+v", fromBlockNum.Int64(), toBlockNum.Int64(), err)
		return err
	}
	for _, l := range sub {
		switch l.Topics[0].String() {
		case util.EventSignHash(chain.TRANSFERTOPIC):
			msg := ErrMsg{
				tp:   el.tokenType,
				from: big.NewInt(int64(l.BlockNumber)),
				to:   big.NewInt(int64(l.BlockNumber)),
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), l.TxHash)
			if err != nil {
				el.errorHandler <- msg
				log.Error("nft TransactionReceipt err : ", err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", l.BlockNumber, err)
				break
			}

			fromAddr := common.HexToAddress(l.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(l.Topics[2].Hex()).String()
			_, txType := el.Accept(fromAddr, toAddr)
			el.Notify(cache.ClearEvent{FromAddr: fromAddr, ToAddr: toAddr})

			_ = model.SpikeTx{
				From:    fromAddr,
				To:      toAddr,
				TxType:  int64(txType),
				TxHash:  l.TxHash.Hex(),
				Status:  int(recp.Status),
				PayTime: int64(block.Time() * 1000),
				TokenId: l.Topics[3].Big().Int64(),
			}
		}
	}
	return nil
}
