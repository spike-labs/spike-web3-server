package sign

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/game"
	"github.com/spike-engine/spike-web3-server/model"
	"github.com/spike-engine/spike-web3-server/util"
	"sync"
	"time"
)

var (
	log       = logger.Logger("sign")
	HwManager *HotWalletManager
)

type HotWalletManager struct {
	scheduler *hotWalletScheduler
	gorm      game.TxTracker
	rdb       *redis.Client
	rLK       sync.RWMutex
}

func NewHWManager() *HotWalletManager {
	m := &HotWalletManager{
		scheduler: newScheduler(),
		gorm:      dao.DbAccessor,
		rdb:       cache.RedisClient,
	}

	_, isNil, err := util.GetIntFromRedis(constant.TOKENID, m.rdb)
	if err != nil {
		panic(err)
	}

	if isNil {
		err := util.SetFromRedis(constant.TOKENID, constant.TOKENID_FROM, 0, m.rdb)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < len(config.Cfg.SignWorkers); i++ {
		worker, err := NewAllRoundWorker(config.Cfg.SignWorkers[i])
		if err != nil {
			log.Error("===Spike log:", err)
			panic(err)
		}
		m.AddWorker(worker)
	}
	go m.scheduler.runSchedule()

	return m
}

func (w *HotWalletManager) AddWorker(worker Worker) {
	w.scheduler.workers = append(w.scheduler.workers, worker)
}

func (w *HotWalletManager) BatchMint(orderId string, tokenURI string, cb string) error {
	w.rLK.Lock()
	TokenId, _, err := util.GetIntFromRedis(constant.TOKENID, w.rdb)
	if err != nil {
		return err
	}

	req := &model.BatchMintReq{
		Uuid:     uuid.New().String(),
		TokenURI: tokenURI,
		TokenID:  TokenId,
	}

	err = util.IncrFromRedis(constant.TOKENID, w.rdb)
	if err != nil {
		return err
	}
	w.rLK.Unlock()

	err = w.gorm.SaveTxCb(model.SpikeTx{
		OrderId:         orderId,
		Uuid:            req.Uuid,
		From:            constant.EmptyAddress,
		To:              config.Cfg.Contract.GameVaultAddress,
		Cb:              cb,
		ContractAddress: config.Cfg.Contract.NftContractAddress[0],
		CreateTime:      time.Now().UnixMilli(),
		TokenId:         TokenId,
	})
	if err != nil {
		return err
	}
	w.scheduler.Schedule(req)
	return nil
}

func (w *HotWalletManager) WithdrawToken(orderId string, toAddress string, amount string, contractAddress string, cb string) error {

	if !util.IsValidAddress(toAddress) || !util.IsValidAddress(contractAddress) {
		return errors.New("=== Spike log : address is error")
	}

	req := &model.WithdrawTokenReq{
		Uuid:         uuid.New().String(),
		ToAddress:    common.HexToAddress(toAddress),
		Amount:       amount,
		TokenAddress: common.HexToAddress(contractAddress),
	}

	err := w.gorm.SaveTxCb(model.SpikeTx{
		OrderId:         orderId,
		Uuid:            req.Uuid,
		From:            config.Cfg.Contract.GameVaultAddress,
		To:              toAddress,
		Cb:              cb,
		ContractAddress: contractAddress,
		CreateTime:      time.Now().UnixMilli(),
		Amount:          amount,
	})
	if err != nil {
		return err
	}
	w.scheduler.Schedule(req)
	return nil
}

func (w *HotWalletManager) WithdrawNFT(orderId string, toAddress string, tokenId int64, contractAddress string, cb string) error {

	if !util.IsValidAddress(toAddress) || !util.IsValidAddress(contractAddress) {
		return errors.New("=== Spike log : address is error")
	}

	req := &model.WithdrawNFTReq{
		Uuid:         uuid.New().String(),
		TokenId:      tokenId,
		ToAddress:    common.HexToAddress(toAddress),
		TokenAddress: common.HexToAddress(contractAddress),
	}

	err := w.gorm.SaveTxCb(model.SpikeTx{
		OrderId:         orderId,
		Uuid:            req.Uuid,
		From:            config.Cfg.Contract.GameVaultAddress,
		To:              toAddress,
		Cb:              cb,
		ContractAddress: contractAddress,
		CreateTime:      time.Now().UnixMilli(),
		TokenId:         tokenId,
	})
	if err != nil {
		return err
	}

	w.scheduler.Schedule(req)
	return nil
}
