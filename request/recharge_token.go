package request

import "github.com/spike-engine/spike-web3-server/service/tx"

type RechargeTokenService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	FromAddress     string `form:"from_address" json:"from_address" binding:"required"`
	Amount          string `form:"amount" json:"amount" binding:"required"`
	ContractAddress int    `form:"contract_address" json:"contract_address" binding:"required"`
	TxHash          string `form:"tx_hash" json:"tx_hash" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
}

func (r *RechargeTokenService) RechargeToken(txSrv *tx.TxService) error {
	err := txSrv.RechargeToken(r.OrderId, r.FromAddress, r.Amount, r.ContractAddress, r.TxHash, r.Cb)
	return err
}
