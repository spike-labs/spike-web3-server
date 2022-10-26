package model

type ApiKey struct {
	Id     string `json:"id" gorm:"primaryKey;column:id;type:varchar(200)"`
	ApiKey string `json:"api_key" gorm:"column:api_key;type:varchar(200)"`
}

func (ApiKey) TableName() string {
	return "api_key"
}
