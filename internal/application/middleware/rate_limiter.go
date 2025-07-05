package middleware

import (
	"log"
	"net"
	"net/http"

	"github.com/sunnygosdk/rate-limiter/internal/application/service"
)

// getRateLimiterByAPIKey returns the rate limiter for the given API key
func getRateLimiterByAPIKey(rateLimiter *service.CacheRateLimiter, w http.ResponseWriter, r *http.Request) bool {
	apiKey := r.Header.Get("API_KEY")

	log.Println("API key", apiKey)
	if apiKey == "" {
		rateLimiter.SetDefaultRateLimiter()
	} else {
		valid := rateLimiter.SetRateLimiterByAPIKey(apiKey)
		if !valid {
			log.Println("Invalid API key", apiKey)
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
			log.Println("Rate limiter not found for API key", r.Header.Get("API_KEY"))
			return
		}

		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("Error getting client IP", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		allowed, err := rateLimiter.Allow(clientIP)
		if err != nil {
			log.Println("Error checking rate limit", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			log.Println("Too many requests for IP", clientIP)
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
