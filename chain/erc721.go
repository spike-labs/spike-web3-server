package chain

import (
	"container/list"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spike-engine/spike-web3-server/cache"
	chain "github.com/spike-engine/spike-web3-server/chain/abi"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/game"
	"github.com/spike-engine/spike-web3-server/util"
	"math/big"
	"sync"
)

func (e *ERC721Listener) Notify(event interface{}) {
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

	e.observers.PushBack(observer)
}

func (e *ERC721Listener) Accept(fromAddr, toAddr string) bool {
	return true
}

type ERC721Listener struct {
	contractAddr   string
	newBlockNotify util.DataChannel
	ec             *ethclient.Client
	abi            abi.ABI
	errorHandler   chan ErrMsg
	observers      *list.List
	observerLk     sync.Mutex
}

func newERC721Listener(contractAddr string, ec *ethclient.Client, newBlockNotify util.DataChannel, abi abi.ABI, errorHandler chan ErrMsg) *ERC721Listener {
	e := &ERC721Listener{
		contractAddr:   contractAddr,
		newBlockNotify: newBlockNotify,
		ec:             ec,
		abi:            abi,
		errorHandler:   errorHandler,
		observers:      list.New(),
	}
	e.AttachObserver(cache.NewManager(cache.RedisClient))
	e.AttachObserver(game.NewCbManager(dao.DbAccessor))
	e.AttachObserver(game.NewNftManager(dao.DbAccessor))
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

	contractAddress := common.HexToAddress(el.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := el.ec.FilterLogs(context.Background(), query)
	if err != nil {
		el.errorHandler <- ErrMsg{
			contractAddr: el.contractAddr,
			from:         fromBlockNum,
			to:           toBlockNum,
		}
		log.Errorf("nft subscribe event log, from: %d,to: %d,err : %+v", fromBlockNum.Int64(), toBlockNum.Int64(), err)
		return err
	}
	for _, logEvent := range sub {
		switch logEvent.Topics[0].String() {
		case util.EventSignHash(chain.TRANSFERTOPIC):
			msg := ErrMsg{
				contractAddr: el.contractAddr,
				from:         big.NewInt(int64(logEvent.BlockNumber)),
				to:           big.NewInt(int64(logEvent.BlockNumber)),
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), logEvent.TxHash)
			if err != nil {
				el.errorHandler <- msg
				log.Error("nft TransactionReceipt err : ", err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(logEvent.BlockNumber)))
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", logEvent.BlockNumber, err)
				break
			}

			fromAddr := common.HexToAddress(logEvent.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(logEvent.Topics[2].Hex()).String()

			go el.Notify(cache.ClearEvent{FromAddr: fromAddr, ToAddr: toAddr})
			go el.Notify(game.NotifyEvent{TxHash: logEvent.TxHash.Hex(), Status: int(recp.Status), PayTime: int64(block.Time() * 1000)})
			go el.Notify(game.NftOwnerUpdateEvent{OwnerAddr: toAddr, ContractAddr: el.contractAddr, TokenId: logEvent.Topics[3].Big().Int64(), UpdateTime: int64(block.Time())})
		}
	}
	return nil
}
