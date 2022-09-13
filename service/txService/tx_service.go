package txService

import (
	"spike-frame/global"
	"spike-frame/model"
	"spike-frame/request"
	"time"
)

type TxService struct {
}

var TxSrv = new(TxService)

func (t *TxService) RechargeToken(service request.RechargeTokenService) error {
	tx := model.SpikeTx{
		OrderId:         service.OrderId,
		From:            service.FromAddress,
		Amount:          service.Amount,
		ContractAddress: service.ContractAddress,
		TxHash:          service.TxHash,
		Cb:              service.Cb,
		CreateTime:      int64(time.Now().Nanosecond()),
	}
	return global.DbAccessor.SaveTxCb(tx)
}

func (t *TxService) ImportNft(service request.ImportNftService) error {
	tx := model.SpikeTx{
		OrderId:         service.OrderId,
		From:            service.From,
		TokenId:         service.TokenId,
		ContractAddress: service.ContractAddress,
		TxHash:          service.TxHash,
		Cb:              service.Cb,
		CreateTime:      int64(time.Now().Nanosecond()),
	}
	return global.DbAccessor.SaveTxCb(tx)
}
