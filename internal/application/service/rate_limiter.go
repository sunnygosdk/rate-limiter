package service

import (
	"context"
	"time"

	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"
)

type CacheRateLimiter struct {
	client  *persistence.CacheClient
	limit   int64
	window  time.Duration
	context context.Context
}

func NewCacheRateLimiter(client *persistence.CacheClient, limit int64, window time.Duration) *CacheRateLimiter {
	return &CacheRateLimiter{
		client:  client,
		limit:   limit,
		window:  window,
		context: context.Background(),
	}
}
