package signService

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	logger "github.com/ipfs/go-log"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/dao"
	"spike-frame/model"
	"spike-frame/request"
	"spike-frame/util"
)

var log = logger.Logger("sign")

type HotWalletManager struct {
	scheduler *hotWalletScheduler
	gorm      *dao.GormAccessor
	rdb       *redis.Client
}

func NewHWManager() (*HotWalletManager, error) {
	m := &HotWalletManager{
		scheduler: newScheduler(),
		gorm:      dao.NewGormAccessor(constant.GormClient),
		rdb:       constant.RedisClient,
	}

	_, isNil, err := util.GetIntFromRedis(constant.TOKENID, m.rdb)
	if err != nil {
		return nil, err
	}

	if isNil {
		util.SetFromRedis(constant.TOKENID, constant.TOKENID_FROM, 0, m.rdb)
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
	err = w.rdb.Incr(context.Background(), constant.TOKENID).Err()
	if err != nil {
		return err
	}

	//err = w.gorm.SaveTxCb(service.OrderId, req.Uuid, "", constant.EmptyAddress, config.Cfg.Contract.GameVaultAddress, "", 0, service.Cb, constant.GAMENFT_IMPORT, time.Now().UnixNano())
	//if err != nil {
	//	return err
	//}
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

	//err := w.gorm.SaveTxCb(service.OrderId, req.Uuid, "", config.Cfg.Contract.GameVaultAddress, service.ToAddress, service.Amount, 0, service.Cb, constant.GAMETOKEN_WITHDRAW, time.Now().UnixNano())
	//if err != nil {
	//	return err
	//}
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

	//err := w.gorm.SaveTxCb(service.OrderId, req.Uuid, "", config.Cfg.Contract.GameVaultAddress, service.ToAddress, "", service.TokenID, service.Cb, constant.GAMENFT_TRANSFER, time.Now().UnixNano())
	//if err != nil {
	//	return err
	//}

	w.scheduler.Schedule(req)
	return nil
}
