package persistence

import (
	"context"
	"time"
)

type CacheClient interface {
	CloseCacheClient() error
	CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error)
}

type CacheRateLimiter struct {
	client  *CacheClient
	limit   int64
	window  time.Duration
	context context.Context
}

func NewCacheRateLimiter(client *CacheClient, limit int64, window time.Duration) *CacheRateLimiter {
	return &CacheRateLimiter{
		client:  client,
		limit:   limit,
		window:  window,
		context: context.Background(),
	}
}
