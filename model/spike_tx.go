package model

type SpikeTx struct {
	OrderId      string `json:"order_id" gorm:"primaryKey;column:order_id;type:varchar(200)"`
	Uuid         string `json:"uuid" gorm:"column:uuid;type:varchar(200)"`
	From         string `json:"from" gorm:"column:from;type:varchar(200)"`
	To           string `json:"to" gorm:"column:to;type:varchar(200)"`
	TxType       int64  `json:"tx_type" gorm:"column:tx_type;type:int"`
	TxHash       string `json:"tx_hash" gorm:"column:tx_hash;type:varchar(200)"`
	Status       int    `json:"status" gorm:"column:status;type:int"`
	NotifyStatus int64  `json:"notify_status" gorm:"column:notify_status;type:int"`
	CreateTime   int64  `json:"create_time" gorm:"column:create_time;comment:game order time;type:int"`
	PayTime      int64  `json:"pay_time" gorm:"column:pay_time;comment:blockchain tx time;type:int"`
	Amount       string `json:"amount" gorm:"column:amount;comment:erc20 tx amount;type:varchar(200)"`
	TokenId      int64  `json:"token_id" gorm:"column:token_id;comment:erc721 tx tokenId;type:bigint"`
	Cb           string `json:"cb" gorm:"column:cb;comment:game cb;type:varchar(200)"`
}

func (SpikeTx) TableName() string {
	return "order"
}
