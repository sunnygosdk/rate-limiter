package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	APP_ENV        string
	APP_PORT       string
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       string
	DEFAULT_LIMIT  int64
	DEFAULT_WINDOW int64
	ADMIN_LIMIT    int64
	ADMIN_WINDOW   int64
	ADMIN_API_KEY  string
	TESTER_LIMIT   int64
	TESTER_WINDOW  int64
	TESTER_API_KEY string
}

func LoadEnvConfig() *EnvConfig {
	env := os.Getenv("APP_ENV")
	if env != "production" {

		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting working directory: %v", err)
		}

		envPath := filepath.Join(wd, ".env")
		err = godotenv.Load(envPath)
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		log.Printf(".env loaded from: %s", envPath)
	}

	return &EnvConfig{
		APP_ENV:        env,
		APP_PORT:       os.Getenv("APP_PORT"),
		REDIS_HOST:     os.Getenv("REDIS_HOST"),
		REDIS_PORT:     os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		REDIS_DB:       os.Getenv("REDIS_DB"),
		DEFAULT_LIMIT:  parseInt64(os.Getenv("DEFAULT_LIMIT"), 10),
		DEFAULT_WINDOW: parseInt64(os.Getenv("DEFAULT_WINDOW"), 60),
		ADMIN_LIMIT:    parseInt64(os.Getenv("ADMIN_LIMIT"), 100),
		ADMIN_WINDOW:   parseInt64(os.Getenv("ADMIN_WINDOW"), 60),
		ADMIN_API_KEY:  os.Getenv("ADMIN_API_KEY"),
		TESTER_LIMIT:   parseInt64(os.Getenv("TESTER_LIMIT"), 50),
		TESTER_WINDOW:  parseInt64(os.Getenv("TESTER_WINDOW"), 60),
		TESTER_API_KEY: os.Getenv("TESTER_API_KEY"),
	}
}

func parseInt64(key string, fallback int64) int64 {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}

	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		log.Printf("Warning: failed to parse %s='%s', using fallback %d", key, valStr, fallback)
		return fallback
	}
	return val
}

func GetEnvConfig() *EnvConfig {
	return LoadEnvConfig()
}
