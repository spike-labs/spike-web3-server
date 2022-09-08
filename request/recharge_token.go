package request

type RechargeTokenService struct {
	OrderId string `json:"order_id" binding:"required"`
	From    string `json:"from" binding:"required"`
	Amount  string `json:"amount" binding:"required"`
	TxType  int64  `json:"tx_type" binding:"required"`
	TxHash  string `json:"tx_hash" binding:"required"`
	Cb      string `json:"cb" binding:"required"`
}
