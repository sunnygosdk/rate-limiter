package middleware

import (
	"net"
	"net/http"

	"github.com/sunnygosdk/rate-limiter/internal/application/service"
)

// getRateLimiterByAPIKey returns the rate limiter for the given API key
func getRateLimiterByAPIKey(rateLimiter *service.CacheRateLimiter, w http.ResponseWriter, r *http.Request) bool {
	apiKey := r.Header.Get("API_KEY")

	if apiKey == "" {
		rateLimiter.SetDefaultRateLimiter()
	} else {
		valid := rateLimiter.SetRateLimiterByAPIKey(apiKey)
		if !valid {
			http.Error(w, "Invalid API_KEY", http.StatusUnauthorized)
			return false
		}
	}

	return true
}

// RateLimiterMiddleware returns a middleware that limits the number of requests per client IP address
func RateLimiterMiddleware(rateLimiter service.CacheRateLimiterClient, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateLimiter := service.NewCacheRateLimiter(rateLimiter)

		if !getRateLimiterByAPIKey(rateLimiter, w, r) {
			return
		}

		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !rateLimiter.Allow(clientIP) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
