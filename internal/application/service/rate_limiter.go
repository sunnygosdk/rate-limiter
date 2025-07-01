package service

import "github.com/sunnygosdk/rate-limiter/internal/infrastructure/persistence"

type CacheClient struct {
	client *persistence.CacheClient
}

func NewCacheClient(client *persistence.CacheClient) *CacheClient {
	return &CacheClient{
		client: client,
	}
}
