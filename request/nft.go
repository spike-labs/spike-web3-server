package request

type NftListService struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}
