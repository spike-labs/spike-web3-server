package dao

import (
	"spike-frame/constant"
)

type UserInfo struct {
	UserName string `json:"username" gorm:"primaryKey;column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserInfo) TableName() string {
	return "user"
}

type UserService struct {
}

var UserSrv = new(UserService)

func (u *UserService) GetUser() (err error, user []UserInfo) {
	var userInfo []UserInfo
	err = constant.GormClient.Select("username", "password").Find(&userInfo).Error
	log.Info("err : ", err)
	return err, userInfo
}
