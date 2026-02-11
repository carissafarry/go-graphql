package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	goredis "github.com/redis/go-redis/v9"
	// redisinfra "go-graphql/internal/infra/cache/redis"
)

func newTestRedisClient() *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
		DB: 9, // NOTE: For testing purposes
	})
}

func TestRedisCacheCRUD(t *testing.T) {
	client := newTestRedisClient()
	cache := NewRedisCache(client)

	ctx := context.Background()
	key := "test:cache:setget" + t.Name()

	t.Cleanup(func() {
		_ = cache.Delete(context.Background(), key)
	})

	// Test Set
	err := cache.Set(ctx, key, []byte("ini set"), 1*time.Minute)
	require.NoError(t, err)

	// // Test Get
	val, err := cache.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, "ini set", string(val))

	// Test Delete
	err = cache.Delete(ctx, key)
	require.NoError(t, err)
}

func TestRedisCacheTTLExpire(t *testing.T) {
	client := newTestRedisClient()
	cache := NewRedisCache(client)

	ctx := context.Background()
	key := "test:cache:ttl:" + t.Name()

	t.Cleanup(func() {
		_ = cache.Delete(context.Background(), key)
	})

	err := cache.Set(ctx, key, []byte("ttl"), 1 * time.Minute)
	require.NoError(t, err)

	// wait 30s
	time.Sleep(30 * time.Second)

	// exists
	_, err = cache.Get(ctx, key)
	require.NoError(t, err)

	// wait ttl expire
	time.Sleep(31 * time.Second)

	// not exists
	_, err = cache.Get(ctx, key)
	require.Error(t, err)
}

// Missing Key should return an error
func TestRedisCacheGetMissingKey(t *testing.T) {
	client := newTestRedisClient()
	cache := NewRedisCache(client)

	ctx := context.Background()
	key := "test:cache:missing:" + t.Name()

	_, err := cache.Get(ctx, key)
	require.Error(t, err)
}

// Overwrite existing key
func TestRedisCacheOverwrite(t *testing.T) {
	client := newTestRedisClient()
	cache := NewRedisCache(client)

	ctx := context.Background()
	key := "test:cache:overwrite:" + t.Name()

	err := cache.Set(ctx, key, []byte("v1"), time.Minute)
	require.NoError(t, err)

	err = cache.Set(ctx, key, []byte("v2"), time.Minute)
	require.NoError(t, err)

	val, err := cache.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, "v2", string(val))
}

// Use default TTL via Options, check if ttl <= 0
func TestRedisCacheDefaultTTL(t *testing.T) {
	client := newTestRedisClient()

	cache := NewRedisCache(client, Options{
		DefaultTTL: 1 * time.Second,
	})

	ctx := context.Background()
	key := "test:cache:defaultttl:" + t.Name()

	t.Cleanup(func() {
		_ = cache.Delete(context.Background(), key)
	})

	// Case 1: Use default TTL (ttl <= 0)
	err := cache.Set(ctx, key, []byte("default"), 0)
	require.NoError(t, err)

	ttl, err := client.TTL(ctx, key).Result()
	require.NoError(t, err)
	require.True(t, ttl > 0 && ttl <= time.Second)
}

// Set with no TTL (ttl = 0), should not expire
func TestRedisCacheWithoutTTL(t *testing.T) {
	client := newTestRedisClient()
	cache := NewRedisCache(client)

	ctx := context.Background()
	key := "test:cache:nottl:" + t.Name()

	t.Cleanup(func() {
		_ = cache.Delete(context.Background(), key)
	})

	err := cache.Set(ctx, key, []byte("without TTL"), 0)
	require.NoError(t, err)

	ttl, err := client.TTL(ctx, key).Result()
	require.NoError(t, err)
	require.Equal(t, time.Duration(-1), ttl) // -1 = no expiry
}

func TestRedisCacheNegativeTTLByDefault(t *testing.T) {
	client := newTestRedisClient()

	cache := NewRedisCache(client, Options{
		DefaultTTL: 2 * time.Second,
	})

	ctx := context.Background()
	key := "test:cache:negativettl:" + t.Name()

	t.Cleanup(func() {
		_ = cache.Delete(context.Background(), key)
	})

	err := cache.Set(ctx, key, []byte("neg"), -1)
	require.NoError(t, err)

	ttl, err := client.TTL(ctx, key).Result()
	require.NoError(t, err)
	require.True(t, ttl > 0 && ttl <= 2*time.Second)
}