package request

type BatchWithdrawalNFTService struct {
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	TokenID         int64  `form:"token_id" json:"token_id" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}
