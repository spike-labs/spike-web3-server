package request

type RechargeTokenService struct {
	OrderId string `form:"order_id" json:"order_id" binding:"required"`
	From    string `form:"from_address" json:"from_address" binding:"required"`
	Amount  string `form:"amount" json:"amount" binding:"required"`
	TxType  int64  `form:"tx_type" json:"tx_type" binding:"required"`
	TxHash  string `form:"tx_hash" json:"tx_hash" binding:"required"`
	Cb      string `form:"cb" json:"cb" binding:"required"`
}
