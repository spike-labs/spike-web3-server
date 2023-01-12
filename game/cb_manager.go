package game

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/model"
	"github.com/spike-engine/spike-web3-server/util"
	"strconv"
	"sync"
	"time"
)

var log = logger.Logger("game")

const (
	HANDLEDURATION          = 10 * time.Minute
	TXTIMEOUTDURATION       = 20 * time.Minute
	LOCKTIMEOUTDURATION     = 5 * time.Minute
	CALLBACKTIMEOUTDURATION = 2 * time.Hour
)

type CallBackService struct {
	TxHash   string `json:"tx_hash"`
	OrderId  string `json:"order_id"`
	TokenId  int    `json:"token_id"`
	TokenUri string `json:"token_uri"`
	Status   int    `json:"status"`
	Sign     string `json:"sign"`
}

var (
	CbMgr     *CbManager
	cbMgrOnce sync.Once
)

type CbManager struct {
	TxTracker
}

type NotifyEvent struct {
	TxHash  string
	Status  int
	PayTime int64
}

func NewCbManager(tracker TxTracker) *CbManager {
	cbMgrOnce.Do(func() {
		CbMgr = &CbManager{tracker}
	})
	return CbMgr
}

func (cm *CbManager) Update(event interface{}) {

	e, ok := event.(NotifyEvent)
	if !ok {
		return
	}
	util.Lock(e.TxHash, constant.TXCBVALUE, LOCKTIMEOUTDURATION, cache.RedisClient)

	defer util.UnLock(e.TxHash, cache.RedisClient)

	txs, err := cm.QueryGameCb(e.TxHash, constant.NOTNOTIFIED)
	if err != nil {
		log.Errorf("query game cb err : %v", err)
		return
	}
	for _, tx := range txs {
		if tx.NotifyStatus == constant.NOTIFIED {
			log.Infof("tx %s has been notify", tx.TxHash)
			continue
		}
		var txStatus int
		txStatus = constant.TXSUCCESS
		if e.Status == 0 {
			txStatus = constant.TXFAILED
		}
		log.Infof("txHash : %s, orderId : %s, cb : %s", tx.TxHash, tx.OrderId, tx.Cb)
		err := cm.UpdateTxStatus(tx.TxHash, txStatus, e.PayTime)
		if err != nil {
			log.Errorf("update tx :%s status err : %v", tx.TxHash, err)
			continue
		}
		err = executeCb(tx, e.Status)
		if err != nil {
			log.Errorf("execute cb order id : %s ,err : %v", tx.OrderId, err)
			continue
		} else {
			err = cm.UpdateTxNotifyStatus(tx.OrderId, constant.NOTIFIED)
			if err != nil {
				log.Errorf("update tx notify status order id : %s ,err : %v", tx.OrderId, err)
			}
		}
	}
}

func (cm *CbManager) Run() {
	ticker := time.NewTicker(HANDLEDURATION)
	for {
		select {
		case <-ticker.C:
			txs, err := cm.QueryNotNotifyTx(constant.NOTNOTIFIED)
			if err != nil {
				log.Infof("query not notify tx err : %v", err)
				break
			}
			for _, tx := range txs {
				cm.handleNotNotifiedTx(tx)
			}
		}
	}
}

func (cm *CbManager) handleNotNotifiedTx(tx model.SpikeTx) {
	if time.Now().After(time.UnixMilli(tx.CreateTime).Add(TXTIMEOUTDURATION)) && time.Now().Before(time.UnixMilli(tx.CreateTime).Add(CALLBACKTIMEOUTDURATION)) {
		util.Lock(tx.TxHash, constant.TXCBVALUE, LOCKTIMEOUTDURATION, cache.RedisClient)

		defer util.UnLock(tx.TxHash, cache.RedisClient)

		log.Infof("tx cb timeout handle, orderId: %s, txHash : %s, createTime : %d", tx.OrderId, tx.TxHash, tx.CreateTime)
		client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
		if err != nil {
			log.Errorf("eth client dial err : %v", err)
			return
		}
		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx.TxHash))
		if err != nil {
			log.Errorf("query tx receipt status err : %v", err)
			return
		}
		var txStatus int
		txStatus = constant.TXSUCCESS
		if receipt.Status == 0 {
			txStatus = constant.TXFAILED
		}
		err = cm.UpdateTxStatus(tx.TxHash, txStatus, tx.PayTime)
		if err != nil {
			log.Errorf("update tx status err : %v", err)
			return
		}
		err = executeCb(tx, int(receipt.Status))
		if err != nil {
			log.Errorf("execute cb order id : %s ,err : %v", tx.OrderId, err)
			return
		}
		err = cm.UpdateTxNotifyStatus(tx.OrderId, constant.NOTIFIED)
		if err != nil {
			log.Errorf("update tx notify status : %s ,err : %v", tx.OrderId, err)
		}
	}
}

func executeCb(tx model.SpikeTx, txStatus int) error {
	isERC721Token := util.IsERC721Token(tx.ContractAddress)
	var signContent string
	var cbReq []byte
	if isERC721Token {
		tokenUri, err := util.QueryNftTokenUri(tx.ContractAddress, string(tx.TokenId))
		if err != nil {
			return err
		}
		signContent = tx.TxHash + tx.OrderId + strconv.FormatInt(int64(txStatus), 10) + strconv.FormatInt(tx.TokenId, 10) + tokenUri
		sign := util.HmacSha256(signContent, config.Cfg.System.SignSecretKey)
		cbReq, _ = json.Marshal(CallBackService{
			TxHash:   tx.TxHash,
			OrderId:  tx.OrderId,
			TokenId:  int(tx.TokenId),
			Status:   txStatus,
			TokenUri: tokenUri,
			Sign:     sign,
		})
		client := resty.New()
		resp, err := client.R().
			SetHeader("Accept", "application/json").
			SetBody(cbReq).
			Post(tx.Cb)
		log.Infof("execute cb : %s, txHash : %s", tx.Cb, tx.TxHash)
		log.Infof("resp : %s", string(resp.Body()))
		return err
	}
	return nil
}
