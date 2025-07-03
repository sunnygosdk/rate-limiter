package main

import (
	"fmt"
	"net/http"

	"github.com/sunnygosdk/rate-limiter/internal/application/middleware"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"
)

func main() {
	AppConfig := config.LoadEnvConfig()

	client := persistence.NewRedisClient(AppConfig)
	defer client.CloseCacheClient()

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	handler := middleware.RateLimiterMiddleware(client, router)
	http.ListenAndServe(fmt.Sprintf(":%s", AppConfig.APP_PORT), handler)
}
