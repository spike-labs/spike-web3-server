package request

type NftListService struct {
	WalletAddress   string `form:"wallet_address" json:"wallet_address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
	Type            string `form:"type" json:"type" binding:"required"`
}

type NftTypeService struct {
	WalletAddress   string `form:"wallet_address" json:"wallet_address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}
