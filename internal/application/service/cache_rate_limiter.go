package service

import (
	"context"

	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"
)

// CacheRateLimiterClient represents a cache rate limiter client
type CacheRateLimiterClient persistence.CacheClient

// CacheRateLimiter represents a cache rate limiter
type CacheRateLimiter struct {
	client      CacheRateLimiterClient
	context     context.Context
	rateLimiter config.RateLimiterConfig
}

// NewCacheRateLimiter creates a new cache rate limiter
func NewCacheRateLimiter(client CacheRateLimiterClient) *CacheRateLimiter {
	return &CacheRateLimiter{
		client:  client,
		context: context.Background(),
	}
}

// SetRateLimiterByAPIKey sets the rate limiter by API key
func (rl *CacheRateLimiter) SetRateLimiterByAPIKey(apiKey string) bool {
	rateLimiter := config.GetRateLimiterByAPIKey(apiKey)
	if rateLimiter == nil {
		return false
	}

	rl.rateLimiter = *rateLimiter
	return true
}

// SetDefaultRateLimiter sets the default rate limiter
func (rl *CacheRateLimiter) SetDefaultRateLimiter() bool {
	rl.rateLimiter = *config.DefaultRateLimiter()
	return true
}

// Allow checks if the request is allowed
func (rl *CacheRateLimiter) Allow(key string) bool {
	count, _ := rl.client.CheckCacheKeysOnWindow(key, rl.context, rl.rateLimiter.Window)
	return count <= rl.rateLimiter.Limit
}
