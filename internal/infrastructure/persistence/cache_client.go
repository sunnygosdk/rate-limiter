package persistence

import (
	"context"
	"log"
	"time"

	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
	"github.com/sunnygosdk/rate-limiter/test/fixture"
)

// CacheClient represents a cache client
type CacheClient interface {
	CloseCacheClient() error
	CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error)
}

// ConfigureCacheClient configures the cache client
func ConfigureCacheClient() CacheClient {
	log.Println("App environment:", config.AppEnvConfig.APP_ENV)
	if config.AppEnvConfig.APP_ENV == "TEST" {
		return fixture.NewRedisClientFixture()
	}

	log.Println("Cache client:", config.AppEnvConfig.CACHE_CLIENT)
	if config.AppEnvConfig.CACHE_CLIENT == "REDIS" {
		return NewRedisClient(config.AppEnvConfig)
	}

	// TODO: Add other cache clients
	log.Println("Cache client not found")
	return nil
}
