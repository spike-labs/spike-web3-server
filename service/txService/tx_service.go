package txService

import (
	"spike-frame/constant"
	"spike-frame/model"
	"spike-frame/request"
	"time"
)

type TxService struct {
}

var TxSrv = new(TxService)

func (t *TxService) RechargeToken(service request.RechargeTokenService) error {
	tx := model.SpikeTx{
		OrderId:    service.OrderId,
		From:       service.From,
		Amount:     service.Amount,
		TxType:     service.TxType,
		TxHash:     service.TxHash,
		Cb:         service.Cb,
		CreateTime: int64(time.Now().Nanosecond()),
	}
	return constant.DbAccessor.SaveTxCb(tx)
}

func (t *TxService) ImportNft(service request.ImportNftService) error {
	tx := model.SpikeTx{
		OrderId:    service.OrderId,
		From:       service.From,
		TokenId:    service.TokenId,
		TxType:     constant.GAMENFT_IMPORT,
		TxHash:     service.TxHash,
		Cb:         service.Cb,
		CreateTime: int64(time.Now().Nanosecond()),
	}
	return constant.DbAccessor.SaveTxCb(tx)
}
