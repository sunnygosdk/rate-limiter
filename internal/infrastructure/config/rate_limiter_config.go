package config

import "time"

var AppConfig *EnvConfig

func init() {
	AppConfig = GetEnvConfig()
}

// RateLimiterConfig represents the configuration for a rate limiter
type RateLimiterConfig struct {
	APIKey string
	Limit  int64
	Window time.Duration
}

// DefaultRateLimiter is the default rate limiter configuration
var DefaultRateLimiter = RateLimiterConfig{
	APIKey: "",
	Limit:  AppConfig.DEFAULT_LIMIT,
	Window: 1 * time.Minute,
}

// AdminRateLimiter is the rate limiter configuration for the admin API key
var AdminRateLimiter = RateLimiterConfig{
	APIKey: "admin",
	Limit:  100,
	Window: 1 * time.Minute,
}

// TesterRateLimiter is the rate limiter configuration for the tester API key
var TesterRateLimiter = RateLimiterConfig{
	APIKey: "tester",
	Limit:  50,
	Window: 1 * time.Minute,
}

// GetRateLimiterByAPIKey returns the rate limiter configuration for the given API key
func GetRateLimiterByAPIKey(apiKey string) *RateLimiterConfig {
	for _, rateLimiter := range []RateLimiterConfig{AdminRateLimiter, TesterRateLimiter} {
		if apiKey == rateLimiter.APIKey {
			return &rateLimiter
		}
	}

	return nil
}
