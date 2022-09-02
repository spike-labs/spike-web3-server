package signService

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"spike-frame/config"
	"spike-frame/model"
	"spike-frame/request"
)

type HotWalletManager struct {
	scheduler *hotWalletScheduler
}

func NewHWManager() (*HotWalletManager, error) {
	m := &HotWalletManager{
		scheduler: newScheduler(),
	}

	for i := 0; i < len(config.Cfg.SignWorkers); i++ {
		worker, err := NewAllRoundWorker()
		if err != nil {
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

func (w *HotWalletManager) BatchMint(service request.BatchMintNFTService) (*model.BatchMintReq, error) {

	req := &model.BatchMintReq{
		Uuid:     uuid.New(),
		TokenID:  service.TokenID,
		TokenURI: service.TokenURI,
	}
	// todo write database
	w.scheduler.Schedule(req)
	return req, nil
}

func (w *HotWalletManager) WithdrawToken(service request.BatchWithdrawalTokenService) (*model.WithdrawTokenReq, error) {
	req := &model.WithdrawTokenReq{
		Uuid:         uuid.New(),
		ToAddress:    common.HexToAddress(service.ToAddress),
		Amount:       service.Amount,
		TokenAddress: common.HexToAddress(service.ContractAddress),
	}
	// todo write database
	w.scheduler.Schedule(req)
	return req, nil
}

func (w *HotWalletManager) WithdrawNFT(service request.BatchWithdrawalNFTService) (*model.WithdrawNFTReq, error) {
	req := &model.WithdrawNFTReq{
		Uuid:         uuid.New(),
		TokenId:      service.TokenID,
		ToAddress:    common.HexToAddress(service.ToAddress),
		TokenAddress: common.HexToAddress(service.ContractAddress),
	}
	// todo write database
	w.scheduler.Schedule(req)
	return req, nil
}
