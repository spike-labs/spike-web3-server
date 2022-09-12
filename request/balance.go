package request

type BalanceService struct {
	WalletAddress string `form:"wallet_address" json:"wallet_address" binding:"required"`
}
