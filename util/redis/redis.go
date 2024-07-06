package redis

import (
	"context"
	"go-boilerplate/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedisClient() (*redis.Client, error) {
	redisConfig := config.Config.Redis

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return RedisClient, nil
}
