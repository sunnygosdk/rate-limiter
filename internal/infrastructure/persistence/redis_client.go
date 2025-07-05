package persistence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
)

// RedisClient represents a Redis client
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a new Redis client
func NewRedisClient(AppConfig *config.EnvConfig) *RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", AppConfig.REDIS_HOST, AppConfig.REDIS_PORT),
		Password: AppConfig.REDIS_PASSWORD,
		DB:       AppConfig.REDIS_DB,
	})
	return &RedisClient{
		client: redisClient,
	}
}

// CheckCacheKeysOnWindow checks the number of requests for a given key in a given window
func (rc *RedisClient) CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error) {
	pipeline := rc.client.Pipeline()
	increment := pipeline.Incr(context, key)
	pipeline.Expire(context, key, window)

	_, err := pipeline.Exec(context)
	if err != nil {
		log.Println("Error checking cache keys on window", err)
		return 0, err
	}

	log.Println("Number of requests for key", key, "in window", window, "is", increment.Val())
	return increment.Val(), nil
}

// CloseCacheClient closes the Redis client
func (rc *RedisClient) CloseCacheClient() error {
	log.Println("Closing Redis client")
	return rc.client.Close()
}
