package persistence

import (
	"context"
	"time"
)

type CacheClient interface {
	CloseCacheClient() error
	CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error)
}
