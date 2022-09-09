package game

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	logger "github.com/ipfs/go-log"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/dao"
	"time"
)

var log = logger.Logger("game")

const (
	handleDuration    = time.Minute
	txTimeoutDuration = 5 * time.Minute
)

type CbManager struct {
	dao.TxTracker
}

type NotifyEvent struct {
	TxHash  string
	Status  int
	PayTime int64
}

func NewCbManager(tracker dao.TxTracker) *CbManager {
	return &CbManager{
		tracker,
	}
}

func (cm *CbManager) Update(event interface{}) {
	if e, ok := event.(NotifyEvent); ok {
		txs, err := cm.QueryGameCb(e.TxHash)
		if err != nil {
			log.Errorf("query game cb err : %v", err)
			return
		}
		for _, tx := range txs {
			if tx.NotifyStatus == constant.NOTIFIED {
				log.Infof("tx %s has been notify", tx.TxHash)
				continue
			}
			log.Infof("cb : %s", tx.Cb)
			err := cm.UpdateTxStatus(tx.TxHash, e.Status, e.PayTime)
			if err != nil {
				log.Errorf("update tx :%s status err : %v", tx.TxHash, err)
				continue
			}
			err = executeCb(tx.Cb)
			if err != nil {
				log.Errorf("execute cb order id : %s ,err : %v", tx.OrderId, err)
				continue
			} else {
				err = cm.UpdateTxNotifyStatus(tx.OrderId, constant.NOTIFIED)
				log.Errorf("update tx notify status order id : %s ,err : %v", tx.OrderId, err)
			}
		}
	}
}

func (cm *CbManager) Run() {
	ticker := time.NewTicker(handleDuration)
	for {
		select {
		case <-ticker.C:
			txs, err := cm.QueryNotNotifyTx()
			if err != nil {
				log.Errorf("query not notify tx err : %v", err)
				break
			}
			for _, tx := range txs {
				if time.Now().After(time.UnixMilli(tx.CreateTime).Add(txTimeoutDuration)) {
					client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
					if err != nil {
						log.Errorf("eth client dial err : %v", err)
						continue
					}
					receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx.TxHash))
					if err != nil {
						log.Errorf("query tx receipt status err : %v", err)
						continue
					}
					err = cm.UpdateTxStatus(tx.TxHash, int(receipt.Status), tx.PayTime)
					if err != nil {
						log.Errorf("update tx status err : %v", err)
						continue
					}
					err = executeCb(tx.Cb)
					notifyStatus := constant.NOTIFIED
					if err != nil {
						notifyStatus = constant.NOTNOTIFIED
						log.Errorf("execute cb order id : %s ,err : %v", tx.OrderId, err)
						continue
					}
					err = cm.UpdateTxNotifyStatus(tx.OrderId, int64(notifyStatus))
					if err != nil {
						log.Errorf("update tx notify status : %s ,err : %v", tx.OrderId, err)
					}
				}
			}
		}
	}
}

func executeCb(url string) error {
	return nil
}
