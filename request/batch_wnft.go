package request

type BatchWithdrawalNFTService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	TokenID         int64  `form:"token_id" json:"token_id" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
}
