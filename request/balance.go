package request

type BalanceService struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}
