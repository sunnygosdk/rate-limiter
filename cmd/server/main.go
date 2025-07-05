package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sunnygosdk/rate-limiter/internal/application/middleware"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"
	"github.com/sunnygosdk/rate-limiter/test/fixture"
)

func main() {
	var client persistence.CacheClient

	log.Println("App environment:", config.AppEnvConfig.APP_ENV)
	if config.AppEnvConfig.APP_ENV == "TEST" {
		log.Println("Using Redis client fixture")
		client = fixture.NewRedisClientFixture()
		defer client.CloseCacheClient()
	} else {
		log.Println("Using Redis client")
		client = persistence.NewRedisClient(config.AppEnvConfig)
		defer client.CloseCacheClient()
	}

	router := GetRouter()
	handler := middleware.RateLimiterMiddleware(client, router)

	log.Printf("Server running on port %s", config.AppEnvConfig.APP_PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", config.AppEnvConfig.APP_PORT), handler)
}
