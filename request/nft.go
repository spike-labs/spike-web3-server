package request

type NftListService struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

type NftTypeService struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}
