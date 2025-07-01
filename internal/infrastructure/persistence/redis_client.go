package persistence

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisClient{
		client: redisClient,
	}
}

func (rc *RedisClient) CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error) {
	pipeline := rc.client.Pipeline()
	increment := pipeline.Incr(context, key)
	pipeline.Expire(context, key, window)

	_, err := pipeline.Exec(context)
	if err != nil {
		return 0, err
	}

	return increment.Val(), nil
}

func (rc *RedisClient) CloseCacheClient() error {
	return rc.client.Close()
}
