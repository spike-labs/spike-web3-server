package constant

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	RedisClient *redis.Client
	GormClient  *gorm.DB
	Viper       *viper.Viper
)
