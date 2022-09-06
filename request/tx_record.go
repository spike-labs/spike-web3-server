package request

type NativeTxRecordService struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type ERC20TxRecordService struct {
	WalletAddress   string `json:"walletAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
}
