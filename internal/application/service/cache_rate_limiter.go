package service

import (
	"context"

	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/shared"
)

type CacheRateLimiterClient persistence.CacheClient

type CacheRateLimiter struct {
	client      CacheRateLimiterClient
	context     context.Context
	rateLimiter shared.RateLimiterConfig
}

func NewCacheRateLimiter(client CacheRateLimiterClient) *CacheRateLimiter {
	return &CacheRateLimiter{
		client:  client,
		context: context.Background(),
	}
}

func (rl *CacheRateLimiter) SetRateLimiterByToken(token string) bool {
	rateLimiter := shared.GetRateLimiterByToken(token)
	if rateLimiter == nil {
		return false
	}

	rl.rateLimiter = *rateLimiter
	return true
}

func (rl *CacheRateLimiter) SetDefaultRateLimiter() bool {
	rl.rateLimiter = shared.DefaultRateLimiter
	return true
}

func (rl *CacheRateLimiter) Allow(key string) bool {
	count, _ := rl.client.CheckCacheKeysOnWindow(key, rl.context, rl.rateLimiter.Window)
	return count <= rl.rateLimiter.Limit
}
