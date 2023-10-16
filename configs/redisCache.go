package configs

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func InitialRedisConnection(config AppConfig, context context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})
	pong, errorConnection := client.Ping(context).Result()
	if errorConnection != nil {
		panic("Error Connect to Redis Cache")
	}
	fmt.Println(pong)

	return client
}
