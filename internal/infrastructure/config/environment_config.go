package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// EnvConfig represents the environment configuration
type EnvConfig struct {
	APP_ENV        string
	APP_PORT       string
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       int
	DEFAULT_LIMIT  int64
	DEFAULT_WINDOW int
	ADMIN_LIMIT    int64
	ADMIN_WINDOW   int
	ADMIN_API_KEY  string
	TESTER_LIMIT   int64
	TESTER_WINDOW  int
	TESTER_API_KEY string
}

var AppEnvConfig *EnvConfig

func init() {
	LoadEnvConfig()
}

// LoadEnvConfig loads the environment configuration
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

	AppEnvConfig = &EnvConfig{
		APP_ENV:        env,
		APP_PORT:       os.Getenv("APP_PORT"),
		REDIS_HOST:     os.Getenv("REDIS_HOST"),
		REDIS_PORT:     os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		REDIS_DB:       parseInt(os.Getenv("REDIS_DB"), 0),
		DEFAULT_LIMIT:  parseInt64(os.Getenv("DEFAULT_LIMIT"), 10),
		DEFAULT_WINDOW: parseInt(os.Getenv("DEFAULT_WINDOW"), 60),
		ADMIN_LIMIT:    parseInt64(os.Getenv("ADMIN_LIMIT"), 100),
		ADMIN_WINDOW:   parseInt(os.Getenv("ADMIN_WINDOW"), 60),
		ADMIN_API_KEY:  os.Getenv("ADMIN_API_KEY"),
		TESTER_LIMIT:   parseInt64(os.Getenv("TESTER_LIMIT"), 50),
		TESTER_WINDOW:  parseInt(os.Getenv("TESTER_WINDOW"), 60),
		TESTER_API_KEY: os.Getenv("TESTER_API_KEY"),
	}

	return AppEnvConfig
}

// parseInt64 parses an environment variable to an int64
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

// parseInt parses an environment variable to an int
func parseInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("Warning: failed to parse %s='%s', using fallback %d", key, valStr, fallback)
		return fallback
	}
	return val
}

// GetEnvConfig returns the environment configuration
func GetEnvConfig() *EnvConfig {
	return AppEnvConfig
}
