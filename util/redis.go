package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

func SetFromRedis(key string, value interface{}, duration time.Duration, client *redis.Client) {
	key = strings.ToLower(key)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if err := client.Set(ctx, key, value, duration).Err(); err != nil {
		log.Errorf("redis set key : %s value %v : err : %v", key, value, err)
	}
}

func GetStringFromRedis(key string, client *redis.Client) (string, error) {
	key = strings.ToLower(key)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		log.Error(err.Error())
		return "", err
	}
	return val, nil
}

func GetIntFromRedis(key string, client *redis.Client) (int64, error) {
	key = strings.ToLower(key)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := client.Get(ctx, key).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		log.Error(err.Error())
		return 0, err
	}
	return val, nil
}

func RmKeyByPrefix(prefix string, client *redis.Client) {
	prefix = strings.ToLower(prefix)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	keys, err := client.Keys(ctx, fmt.Sprintf("%s*", prefix)).Result()
	if err != nil {
		log.Error(err)
	}
	log.Error(client.Del(ctx, keys...).Err())
}
