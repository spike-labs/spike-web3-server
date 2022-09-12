package chain

import (
	"container/list"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"spike-frame/cache"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/game"
	"spike-frame/model"
	"spike-frame/util"
	"strings"
	"sync"
	"time"
)

func (bl *BNBListener) Notify(event interface{}) {
	bl.observerLk.Lock()
	defer bl.observerLk.Unlock()

	for o := bl.observers.Front(); o != nil; o = o.Next() {
		if ov, ok := o.Value.(cache.Observer); ok {
			ov.Update(event)
		}
	}
}

func (bl *BNBListener) AttachObserver(observer cache.Observer) {
	bl.observerLk.Lock()
	defer bl.observerLk.Unlock()

	o := bl.observers.Front()
	if o == nil {
		bl.observers.PushBack(observer)
	}
}

func (bl *BNBListener) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(toAddr) {
		return true, constant.NATIVE_RECHARGE
	}

	return false, constant.NOT_EXIST
}

type BNBListener struct {
	ec           *ethclient.Client
	chainId      *big.Int
	errorHandler chan ErrMsg
	observers    *list.List
	observerLk   sync.Mutex
}

func newBNBListener(ec *ethclient.Client, errorHandler chan ErrMsg) *BNBListener {
	chainId, err := ec.NetworkID(context.Background())
	if err != nil {
		panic("query network id err")
		return nil
	}
	bl := &BNBListener{
		ec:           ec,
		chainId:      chainId,
		errorHandler: errorHandler,
	}
	bl.AttachObserver(game.NewCbManager(constant.DbAccessor))
	return bl
}

func (bl *BNBListener) run() {
	go bl.NewBlockFilter()
}

func (bl *BNBListener) NewBlockFilter() error {
	newBlockChan := make(chan *types.Header)
	sub, err := bl.ec.SubscribeNewHead(context.Background(), newBlockChan)
	if err != nil {
		log.Error("bnb subscribe new head err : ", err)
		return err
	}
	for {
		select {
		case err = <-sub.Err():
			sub = event.Resubscribe(time.Millisecond, func(ctx context.Context) (event.Subscription, error) {
				return bl.ec.SubscribeNewHead(context.Background(), newBlockChan)
			})
			log.Error("new block subscribe err : ", err)
		case header := <-newBlockChan:
			height := new(big.Int).Sub(header.Number, big.NewInt(constant.BlockConfirmHeight))
			cacheHeight, _, err := util.GetIntFromRedis(BLOCKNUM+config.Cfg.System.MachineId, constant.RedisClient)

			if height.Int64()-1 > cacheHeight {
				for i := cacheHeight + 1; i < height.Int64(); i++ {
					log.Infof("ws node timeout err : height %d", i)
					bl.errorHandler <- ErrMsg{
						tp:   model.Bnb,
						from: big.NewInt(i),
						to:   big.NewInt(i),
					}
					util.Eb.Publish(constant.NewBlockTopic, big.NewInt(i))
				}
			}
			util.Eb.Publish(constant.NewBlockTopic, height)
			log.Infof("new block num : %d, height : %d", header.Number.Int64(), height.Int64())

			err = bl.SingleBlockFilter(height)
			if err != nil {
				bl.errorHandler <- ErrMsg{
					tp:   model.Bnb,
					from: height,
					to:   height,
				}
			}
			util.SetFromRedis(BLOCKNUM+config.Cfg.System.MachineId, height.Int64(), 0, constant.RedisClient)
			log.Infof("bnb listen new block %d finished", height)
		}
	}
}

func (bl *BNBListener) handlePastBlock(blockNum, nowBlockNum *big.Int) error {
	throttle := make(chan struct{}, 30)
	var wg sync.WaitGroup
	wg.Add(int(new(big.Int).Sub(nowBlockNum, blockNum).Int64()) + 1)
	for i := blockNum.Int64(); i <= nowBlockNum.Int64(); i++ {
		throttle <- struct{}{}
		go func(height int64) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			h := big.NewInt(height)
			err := bl.SingleBlockFilter(h)
			if err != nil {
				bl.errorHandler <- ErrMsg{
					tp:   model.Bnb,
					from: h,
					to:   h,
				}
			}
		}(i)
	}
	wg.Wait()
	return nil
}

func (bl *BNBListener) SingleBlockFilter(height *big.Int) error {
	block, err := bl.ec.BlockByNumber(context.Background(), height)
	if err != nil {
		log.Errorf("bnb blockByHash heght : %d ,err : %+v", height.Int64(), err)
		return err
	}
	log.Infof("bnb height : %d , tx num :  %d", block.Number(), len(block.Transactions()))
	for _, tx := range block.Transactions() {
		var fromAddr string
		if msg, err := tx.AsMessage(types.NewEIP155Signer(bl.chainId), nil); err == nil {
			fromAddr = msg.From().Hex()
		}
		if tx.To() == nil {
			continue
		}
		if tx.Value().Int64() == 0 {
			continue
		}
		accept, _ := bl.Accept(fromAddr, tx.To().Hex())
		if !accept {
			continue
		}
		recp, err := bl.ec.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Error("bnb TransactionReceipt err : ", err)
			return err
		}
		bl.Notify(game.NotifyEvent{TxHash: tx.Hash().Hex(), Status: int(recp.Status), PayTime: int64(block.Time() * 1000)})
		log.Infof("native tx ,from :%s, to : %s,  amount : %s", fromAddr, tx.To().Hex(), tx.Value().String())
	}
	return nil
}
