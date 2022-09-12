package request

type NftListService struct {
	WalletAddress string `form:"wallet_address" json:"wallet_address" binding:"required"`
	Type          string `form:"type" json:"type" binding:"required"`
}

type NftTypeService struct {
	WalletAddress string `form:"wallet_address" json:"wallet_address" binding:"required"`
}
