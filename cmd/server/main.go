package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sunnygosdk/rate-limiter/internal/application/middleware"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"
)

func main() {
	client := persistence.ConfigureCacheClient()
	if client == nil {
		log.Fatal("Cache client not found")
	}
	defer client.CloseCacheClient()

	router := GetRouter()
	handler := middleware.RateLimiterMiddleware(client, router)

	log.Printf("Server running on port %s", config.AppEnvConfig.APP_PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.AppEnvConfig.APP_PORT), handler)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
