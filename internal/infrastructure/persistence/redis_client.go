package persistence

import (
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

type RedisRateLimiter struct {
	BaseRateLimiter
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
}

func (rc *RedisClient) Allow(key string) bool {
	return true
}

func (rc *RedisClient) CloseCacheClient() error {
	return rc.client.Close()
}
