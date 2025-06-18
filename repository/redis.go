package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.197.132:6379",
		Password: "20031001",
		DB:       1,
	})
}
