package model

type NftOwner struct {
	Id              string `json:"id" gorm:"primaryKey;column:id;type:varchar(200)"`
	OwnerAddress    string `json:"owner_address" gorm:"column:owner_address;type:varchar(200)"`
	ContractAddress string `json:"contract_address" gorm:"column:contract_address;type:varchar(200)"`
	UpdateTime      int64  `json:"update_time" gorm:"column:update_time;comment:nft owner update time;type:int;default:null"`
	TokenId         int64  `json:"token_id" gorm:"column:token_id;comment:erc721 tx tokenId;type:bigint;default:null"`
}

func (NftOwner) TableName() string {
	return "nft_owner"
}
