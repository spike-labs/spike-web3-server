package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"github.com/spike-engine/spike-web3-server/dao"
)

var (
	RedisClient *redis.Client
	GormClient  *gorm.DB
	Viper       *viper.Viper
	DbAccessor  *dao.GormAccessor
)
