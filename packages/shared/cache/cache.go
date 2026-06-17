package cache

import (
	"context"

	"github.com/fintech/marketpilot/packages/shared/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func New(
	cfg *config.Config,
) (*Cache, error) {

	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
	})

	if err := client.Ping(
		context.Background(),
	).Err(); err != nil {
		return nil, err
	}

	return &Cache{
		Client: client,
	}, nil
}

func (c *Cache) Close() error {
	return c.Client.Close()
}