package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/spike-engine/spike-web3-server/dao"
	"gorm.io/gorm"
)

var (
	RedisClient *redis.Client
	GormClient  *gorm.DB
	Viper       *viper.Viper
	DbAccessor  *dao.GormAccessor
)
