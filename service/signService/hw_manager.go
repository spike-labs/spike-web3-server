package signService

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	logger "github.com/ipfs/go-log"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/game"
	"spike-frame/model"
	"spike-frame/request"
	"spike-frame/util"
	"time"
)

var log = logger.Logger("sign")

type HotWalletManager struct {
	scheduler *hotWalletScheduler
	gorm      game.TxTracker
	rdb       *redis.Client
}

func NewHWManager() (*HotWalletManager, error) {
	m := &HotWalletManager{
		scheduler: newScheduler(),
		gorm:      constant.DbAccessor,
		rdb:       constant.RedisClient,
	}

	_, isNil, err := util.GetIntFromRedis(constant.TOKENID, m.rdb)
	if err != nil {
		return nil, err
	}

	if isNil {
		err := util.SetFromRedis(constant.TOKENID, constant.TOKENID_FROM, 0, m.rdb)
		if err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(config.Cfg.SignWorkers); i++ {
		worker, err := NewAllRoundWorker()
		if err != nil {
			log.Error("===Spike log:", err)
			return nil, err
		}
		m.AddWorker(worker)
	}
	go m.scheduler.runSchedule()

	return m, nil
}

func (w *HotWalletManager) AddWorker(worker Worker) {
	w.scheduler.workers = append(w.scheduler.workers, worker)
}

func (w *HotWalletManager) BatchMint(service request.BatchMintNFTService) error {

	TokenId, _, err := util.GetIntFromRedis(constant.TOKENID, w.rdb)
	if err != nil {
		return err
	}

	req := &model.BatchMintReq{
		Uuid:     uuid.New().String(),
		TokenURI: service.TokenURI,
		TokenID:  TokenId,
	}

	err = util.IncrFromRedis(constant.TOKENID, w.rdb)
	if err != nil {
		return err
	}

	err = w.gorm.SaveTxCb(model.SpikeTx{
		OrderId:    service.OrderId,
		Uuid:       req.Uuid,
		From:       constant.EmptyAddress,
		To:         config.Cfg.Contract.GameVaultAddress,
		Cb:         service.Cb,
		TxType:     constant.GAMENFT_TRANSFER,
		CreateTime: time.Now().UnixMilli(),
		TokenId:    TokenId,
	})
	if err != nil {
		return err
	}
	w.scheduler.Schedule(req)
	return nil
}

func (w *HotWalletManager) WithdrawToken(service request.BatchWithdrawalTokenService) error {

	if !util.IsValidAddress(service.ToAddress) || !util.IsValidAddress(service.ContractAddress) {
		return errors.New("=== Spike log : address is error")
	}

	req := &model.WithdrawTokenReq{
		Uuid:         uuid.New().String(),
		ToAddress:    common.HexToAddress(service.ToAddress),
		Amount:       service.Amount,
		TokenAddress: common.HexToAddress(service.ContractAddress),
	}
	var txType int64
	switch service.ContractAddress {
	case config.Cfg.Contract.GovernanceTokenAddress:
		txType = constant.GOVERNANCE_WITHDRAW
	case config.Cfg.Contract.GameTokenAddress:
		txType = constant.GAMETOKEN_WITHDRAW
	case config.Cfg.Contract.UsdcAddress:
		txType = constant.USDC_WITHDRAW
	case constant.EmptyAddress:
		txType = constant.NATIVE_WITHDRAW
	}

	err := w.gorm.SaveTxCb(model.SpikeTx{
		OrderId:    service.OrderId,
		Uuid:       req.Uuid,
		From:       config.Cfg.Contract.GameVaultAddress,
		To:         service.ToAddress,
		Cb:         service.Cb,
		TxType:     txType,
		CreateTime: time.Now().UnixMilli(),
		Amount:     service.Amount,
	})
	if err != nil {
		return err
	}
	w.scheduler.Schedule(req)
	return nil
}

func (w *HotWalletManager) WithdrawNFT(service request.BatchWithdrawalNFTService) error {

	if !util.IsValidAddress(service.ToAddress) || !util.IsValidAddress(service.ContractAddress) {
		return errors.New("=== Spike log : address is error")
	}

	req := &model.WithdrawNFTReq{
		Uuid:         uuid.New().String(),
		TokenId:      service.TokenID,
		ToAddress:    common.HexToAddress(service.ToAddress),
		TokenAddress: common.HexToAddress(service.ContractAddress),
	}

	err := w.gorm.SaveTxCb(model.SpikeTx{
		OrderId:    service.OrderId,
		Uuid:       req.Uuid,
		From:       config.Cfg.Contract.GameVaultAddress,
		To:         service.ToAddress,
		Cb:         service.Cb,
		TxType:     constant.GAMENFT_TRANSFER,
		CreateTime: time.Now().UnixMilli(),
		TokenId:    service.TokenID,
	})
	if err != nil {
		return err
	}

	w.scheduler.Schedule(req)
	return nil
}
