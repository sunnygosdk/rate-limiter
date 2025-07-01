package persistence

import (
	"context"
	"time"
)

type CacheClient interface {
	CloseCacheClient() error
	Allow(key string) bool
}

type BaseRateLimiter struct {
	limit   int64
	window  time.Duration
	context context.Context
}
