package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/repository/interfaces"
)

type redisDB struct {
	client *redis.Client
}

func NewCacheRepo(cfg config.Config) interfaces.CacheRepo {

	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	return &redisDB{
		client: client,
	}
}

func (c *redisDB) Set(ctx context.Context, key string, data []byte, duration time.Duration) error {

	err := c.client.Set(key, data, duration).Err()

	return err
}
func (c *redisDB) Get(ctx context.Context, key string) ([]byte, error) {

	result, err := c.client.Get(key).Result()

	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

func (c *redisDB) Del(ctx context.Context, key string) error {

	return c.client.Del(key).Err()
}
