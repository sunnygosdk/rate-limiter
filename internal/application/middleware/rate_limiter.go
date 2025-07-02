package middleware

import (
	"net/http"

	"github.com/sunnygosdk/rate-limiter/internal/application/service"
)

func RateLimiterMiddleware(rateLimiter service.CacheRateLimiterClient, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateLimiter := service.NewCacheRateLimiter(rateLimiter)

		apiKey := r.Header.Get("API_KEY")
		if apiKey == "" {
			rateLimiter.SetDefaultRateLimiter()
		} else {
			valid := rateLimiter.SetRateLimiterByToken(apiKey)
			if !valid {
				http.Error(w, "Invalid API_KEY", http.StatusUnauthorized)
				return
			}
		}

		if !rateLimiter.Allow(apiKey) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
