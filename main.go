package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	client  *redis.Client
	limit   int64
	window  time.Duration
	context context.Context
}

func NewRateLimiter(client *redis.Client, limit int64, window time.Duration) *RateLimiter {
	return &RateLimiter{
		client:  client,
		limit:   limit,
		window:  window,
		context: context.Background(),
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	pipeline := rl.client.TxPipeline()
	increment := pipeline.Incr(rl.context, key)
	pipeline.Expire(rl.context, key, rl.window)

	_, err := pipeline.Exec(rl.context)
	if err != nil {
		return false
	}

	return increment.Val() <= rl.limit
}

func rateLimiterMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !rl.Allow(clientIP) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer client.Close()

	rateLimiter := NewRateLimiter(client, 10, 1*time.Minute)

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	handler := rateLimiterMiddleware(rateLimiter, router)

	http.ListenAndServe(":8080", handler)
}
