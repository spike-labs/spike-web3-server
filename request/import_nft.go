package request

type ImportNftService struct {
	OrderId string `json:"order_id" binding:"required"`
	From    string `json:"from" binding:"required"`
	TokenId int64  `json:"token_id" binding:"required"`
	TxHash  string `json:"tx_hash" binding:"required"`
	Cb      string `json:"cb" binding:"required"`
}
