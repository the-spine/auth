package db

import (
	"auth/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(config *config.Config) error {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	res, err := RedisClient.Ping(context.Background()).Result()

	log.Printf("Redis Says Ping %s", res)

	return err
}
