package fixture

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sunnygosdk/rate-limiter/internal/infrastructure/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// RedisClientFixture represents a Redis client fixture
type RedisClientFixture struct {
	container testcontainers.Container
	config    *config.EnvConfig
}

// SetupRedisContainer sets up a Redis container
func SetupRedisContainer() (*RedisClientFixture, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:7.2-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	host, err := redisC.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := redisC.MappedPort(ctx, "6379")
	if err != nil {
		return nil, err
	}

	opts := &config.EnvConfig{
		REDIS_HOST:     host,
		REDIS_PORT:     port.Port(),
		REDIS_DB:       0,
		REDIS_PASSWORD: "",
	}

	return &RedisClientFixture{
		container: redisC,
		config:    opts,
	}, nil
}

// NewRedisClientFixture creates a new Redis client fixture
func NewRedisClientFixture() *RedisClientFixture {
	container, err := SetupRedisContainer()
	if err != nil {
		return nil
	}

	return &RedisClientFixture{
		container: container.container,
		config:    container.config,
	}
}

// CloseCacheClient closes the Redis client fixture
func (rc *RedisClientFixture) CloseCacheClient() error {
	return rc.container.Terminate(context.Background())
}

// CheckCacheKeysOnWindow checks the number of requests for a given key in a given window
func (rc *RedisClientFixture) CheckCacheKeysOnWindow(key string, ctx context.Context, window time.Duration) (int64, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rc.config.REDIS_HOST, rc.config.REDIS_PORT),
		Password: rc.config.REDIS_PASSWORD,
		DB:       rc.config.REDIS_DB,
	})
	pipeline := client.Pipeline()
	increment := pipeline.Incr(ctx, key)
	pipeline.Expire(ctx, key, window)

	_, err := pipeline.Exec(ctx)
	if err != nil {
		return 0, err
	}

	return increment.Val(), nil
}
