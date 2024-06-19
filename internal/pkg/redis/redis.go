package redis

import (
	"context"
	"dxj/configs"
	"dxj/internal/pkg/logger"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Connect(config *configs.Redis) *redis.Client {
	logger := logger.LogrusLogger
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.PassWord,
		DB:       config.Db,
	})
	ping := rdb.Ping(ctx)
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	logger.Printf(`🍟: Successfully connected to Redis at ` + address)
	return rdb
}
