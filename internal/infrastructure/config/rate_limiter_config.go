package config

import "time"

// RateLimiterConfig represents the configuration for a rate limiter
type RateLimiterConfig struct {
	APIKey string
	Limit  int64
	Window time.Duration
}

// DefaultRateLimiter is the default rate limiter configuration
func DefaultRateLimiter() *RateLimiterConfig {
	return &RateLimiterConfig{
		APIKey: "",
		Limit:  AppEnvConfig.DEFAULT_LIMIT,
		Window: time.Duration(AppEnvConfig.DEFAULT_WINDOW) * time.Second,
	}
}

// AdminRateLimiter is the rate limiter configuration for the admin API key
func AdminRateLimiter() *RateLimiterConfig {
	return &RateLimiterConfig{
		APIKey: AppEnvConfig.ADMIN_API_KEY,
		Limit:  AppEnvConfig.ADMIN_LIMIT,
		Window: time.Duration(AppEnvConfig.ADMIN_WINDOW) * time.Second,
	}
}

// TesterRateLimiter is the rate limiter configuration for the tester API key
func TesterRateLimiter() *RateLimiterConfig {
	return &RateLimiterConfig{
		APIKey: AppEnvConfig.TESTER_API_KEY,
		Limit:  AppEnvConfig.TESTER_LIMIT,
		Window: time.Duration(AppEnvConfig.TESTER_WINDOW) * time.Second,
	}
}

// GetRateLimiterByAPIKey returns the rate limiter configuration for the given API key
func GetRateLimiterByAPIKey(apiKey string) *RateLimiterConfig {
	for _, rateLimiter := range []*RateLimiterConfig{AdminRateLimiter(), TesterRateLimiter()} {
		if apiKey == rateLimiter.APIKey {
			return rateLimiter
		}
	}

	return nil
}
