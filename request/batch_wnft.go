package request

import "github.com/spike-engine/spike-web3-server/service/sign"

type BatchWithdrawalNFTService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	TokenID         int64  `form:"token_id" json:"token_id" binding:"required"`
	ContractAddress int    `form:"contract_address" json:"contract_address" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
}

func (b *BatchWithdrawalNFTService) WithdrawNFT(hwManager *sign.HotWalletManager) error {
	err := hwManager.WithdrawNFT(b.OrderId, b.ToAddress, b.TokenID, b.ContractAddress, b.Cb)
	return err
}
