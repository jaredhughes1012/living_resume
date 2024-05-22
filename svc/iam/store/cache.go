package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	// Caches an activation code for a requested account
	CacheActivationCode(ctx context.Context, code, email string) error

	// Gets a cached email address for an activation code
	GetActivationEmail(ctx context.Context, code string) (string, error)
}

type redisCache struct {
	rdb *redis.Client
}

// GetActivationEmail implements Cache.
func (r *redisCache) GetActivationEmail(ctx context.Context, code string) (string, error) {
	return r.rdb.Get(ctx, code).Result()
}

// CacheActivationCode implements Cache.
func (r *redisCache) CacheActivationCode(ctx context.Context, code, email string) error {
	return r.rdb.Set(ctx, code, email, time.Hour*1).Err()
}

func NewCache(ctx context.Context, connStr string) (Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: connStr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &redisCache{rdb: rdb}, nil
}
