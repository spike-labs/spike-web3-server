package request

import "github.com/spike-engine/spike-web3-server/service/sign"

type BatchWithdrawalTokenService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	Amount          string `form:"amount" json:"amount" binding:"required"`
	ContractAddress int    `form:"contract_address" json:"contract_address"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
}

func (b *BatchWithdrawalTokenService) WithdrawToken(hwManager *sign.HotWalletManager) error {
	err := hwManager.WithdrawToken(b.OrderId, b.ToAddress, b.Amount, b.ContractAddress, b.Cb)
	return err
}
