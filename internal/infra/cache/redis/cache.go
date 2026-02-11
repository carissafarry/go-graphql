package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
    Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
    Get(ctx context.Context, key string) ([]byte, error)
    Delete(ctx context.Context, key string) error
}

type RedisCache struct {
    client *redis.Client
	defaultTTL time.Duration
}

type Options struct {
	DefaultTTL time.Duration
}

func NewRedisCache(
	client *redis.Client,
	opts ...Options,
) *RedisCache {
    cache := &RedisCache{client: client}

	if len(opts) > 0 {
		cache.defaultTTL = opts[0].DefaultTTL
	}
	return cache
}

func (r *RedisCache) Set(ctx context.Context, key string, val []byte, ttl time.Duration) error {
	if ttl <= 0{
		ttl = r.defaultTTL
	}
	if ttl > 0 {
		return r.client.SetEx(ctx, key, val, ttl).Err()
	}
	return r.client.Set(ctx, key, val, 0).Err() // if ttl < 0, no expiration
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, key).Err()
}