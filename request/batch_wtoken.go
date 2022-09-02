package request

type BatchWithdrawalTokenService struct {
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	Amount          int64  `form:"amount" json:"amount" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}
