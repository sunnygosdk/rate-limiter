package persistence

import (
	"context"
	"time"
)

// CacheClient represents a cache client
type CacheClient interface {
	CloseCacheClient() error
	CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error)
}
