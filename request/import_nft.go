package request

type ImportNftService struct {
	OrderId string `form:"order_id" json:"order_id" binding:"required"`
	From    string `form:"from_address" json:"from_address" binding:"required"`
	TokenId int64  `form:"token_id" json:"token_id" binding:"required"`
	TxHash  string `form:"tx_hash" json:"tx_hash" binding:"required"`
	Cb      string `form:"cb" json:"cb" binding:"required"`
}
