package request

type NativeTxRecordService struct {
	WalletAddress string `form:"wallet_address" json:"wallet_address" binding:"required"`
}

type ERC20TxRecordService struct {
	WalletAddress   string `form:"wallet_address" json:"wallet_address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}
