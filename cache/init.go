package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	logger "github.com/ipfs/go-log"
	"spike-frame/config"
)

var log = logger.Logger("cache")

func ConnectRedis() *redis.Client {
	m := config.Cfg.Redis

	client := redis.NewClient(
		&redis.Options{
			Addr:       m.Address,
			Password:   m.Password,
			MaxRetries: 1,
		})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Error("redis init err : ", err)
		panic("redis error")
		return nil
	}
	return client
}
