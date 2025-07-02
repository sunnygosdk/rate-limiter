package shared

import "time"

type RateLimiterConfig struct {
	Token  string
	Limit  int64
	Window time.Duration
}

var DefaultRateLimiter = RateLimiterConfig{
	Token:  "default",
	Limit:  10,
	Window: 1 * time.Minute,
}

var AdminRateLimiter = RateLimiterConfig{
	Token:  "admin",
	Limit:  100,
	Window: 1 * time.Minute,
}

var TesterRateLimiter = RateLimiterConfig{
	Token:  "tester",
	Limit:  50,
	Window: 1 * time.Minute,
}

func GetRateLimiterByToken(token string) *RateLimiterConfig {
	for _, rateLimiter := range []RateLimiterConfig{AdminRateLimiter, TesterRateLimiter} {
		if token == rateLimiter.Token {
			return &rateLimiter
		}
	}

	return nil
}
