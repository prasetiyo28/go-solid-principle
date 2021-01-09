package redis

import (
	"github.com/go-redis/redis"
)

func RedisInit() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	return redisClient

}
