package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/dao"
)

var (
	log         = logger.Logger("cache")
	RedisClient *redis.Client
)

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
	apiKeyList, err := dao.DbAccessor.QueryApiKey()
	for _, apiKey := range apiKeyList {
		client.SAdd(context.Background(), "api_key", apiKey.ApiKey)
	}
	return client
}
