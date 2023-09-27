package interfaces

import (
	"context"
	"time"
)

type CacheRepo interface {
	Set(ctx context.Context, key string, data []byte, duration time.Duration) error
	Get(ctx context.Context, key string) (data []byte, err error)
	Del(ctx context.Context, key string) error
}
