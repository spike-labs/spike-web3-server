package chain

import (
	"container/list"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"github.com/spike-engine/spike-web3-server/cache"
	chain "github.com/spike-engine/spike-web3-server/chain/abi"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/game"
	"github.com/spike-engine/spike-web3-server/global"
	"github.com/spike-engine/spike-web3-server/util"
	"strings"
	"sync"
)

func (e *ERC20Listener) Notify(event interface{}) {
	e.observerLk.Lock()
	defer e.observerLk.Unlock()

	for o := e.observers.Front(); o != nil; o = o.Next() {
		if ov, ok := o.Value.(cache.Observer); ok {
			ov.Update(event)
		}
	}
}

func (e *ERC20Listener) AttachObserver(observer cache.Observer) {
	e.observerLk.Lock()
	defer e.observerLk.Unlock()

	o := e.observers.Front()
	if o == nil {
		e.observers.PushBack(observer)
	}
}

func (e *ERC20Listener) Accept(fromAddr, toAddr string) bool {
	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(toAddr) {
		return true
	}

	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(fromAddr) {
		return true
	}

	return false
}

type ERC20Listener struct {
	contractAddr   string
	newBlockNotify util.DataChannel
	ec             *ethclient.Client
	abi            abi.ABI
	errorHandler   chan ErrMsg
	observers      *list.List
	observerLk     sync.Mutex
}

func newERC20Listener(contractAddr string, ec *ethclient.Client, newBlockNotify util.DataChannel, abi abi.ABI, errorHandler chan ErrMsg) *ERC20Listener {
	el := &ERC20Listener{
		contractAddr:   contractAddr,
		newBlockNotify: newBlockNotify,
		ec:             ec,
		abi:            abi,
		errorHandler:   errorHandler,
		observers:      list.New(),
	}
	el.AttachObserver(game.NewCbManager(global.DbAccessor))
	return el
}

func (el *ERC20Listener) run() {
	go el.NewEventFilter(el.contractAddr)
}

func (el *ERC20Listener) NewEventFilter(contractAddr string) error {
	for {
		select {
		case de := <-el.newBlockNotify:
			height := de.Data.(*big.Int)
			el.handlePastBlock(height, height)
		}
	}
}

func (el *ERC20Listener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("erc20 past event filter, type : %v, fromBlock : %d, toBlock : %d ", el.contractAddr, fromBlockNum, toBlockNum)
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
			contractAddr: el.contractAddr,
			from:         fromBlockNum,
			to:           toBlockNum,
		}
		log.Errorf("erc20 subscribe err : %+v, from : %d, to : %d, type : %s", err, fromBlockNum.Int64(), toBlockNum.Int64(), el.contractAddr)
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

			input, err := el.abi.Events["Transfer"].Inputs.Unpack(logEvent.Data)
			if err != nil {
				log.Error("erc20 data unpack err : ", err)
				el.errorHandler <- msg
				break
			}
			fromAddr := common.HexToAddress(logEvent.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(logEvent.Topics[2].Hex()).String()
			if accept := el.Accept(fromAddr, toAddr); !accept {
				break
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), logEvent.TxHash)
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query txReceipt txHash : %s, err : %+v", logEvent.TxHash, err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(logEvent.BlockNumber)))
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", logEvent.BlockNumber, err)
				break
			}
			log.Infof("erc20 tx ,from :%s, to : %s, type : %s,  amount : %s", fromAddr, toAddr, el.contractAddr, input[0].(*big.Int).String())
			go el.Notify(game.NotifyEvent{TxHash: logEvent.TxHash.Hex(), Status: int(recp.Status), PayTime: int64(block.Time() * 1000)})
		case util.EventSignHash(chain.WITHRAWALTOPIC):
			msg := ErrMsg{
				contractAddr: el.contractAddr,
				from:         big.NewInt(int64(logEvent.BlockNumber)),
				to:           big.NewInt(int64(logEvent.BlockNumber)),
			}
			input, err := el.abi.Events["Withdraw"].Inputs.Unpack(logEvent.Data)
			if err != nil {
				log.Error("game vault data unpack err : ", err)
				el.errorHandler <- msg
				break
			}
			if input[0].(common.Address).String() != constant.EmptyAddress {
				break
			}
			fromAddr := input[1].(common.Address).String()
			toAddr := input[2].(common.Address).String()
			if accept := el.Accept(fromAddr, toAddr); !accept {
				break
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), logEvent.TxHash)
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query txReceipt txHash : %s, err : %+v", logEvent.TxHash, err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(logEvent.BlockNumber)))
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", logEvent.BlockNumber, err)
				break
			}
			log.Infof("erc20 tx ,from :%s, to : %s, type : %s,  amount : %s", fromAddr, toAddr, el.contractAddr, input[3].(*big.Int).String())
			go el.Notify(game.NotifyEvent{TxHash: logEvent.TxHash.Hex(), Status: int(recp.Status), PayTime: int64(block.Time() * 1000)})
		}
	}
	return err
}
