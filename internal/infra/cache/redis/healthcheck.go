package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Healthcheck checks whether Redis is reachable and ready to serve requests.
func Healthcheck(ctx context.Context, client *goredis.Client) error {
	if client == nil {
		return fmt.Errorf("redis healthcheck failed: redis client is nil")
	}

	// Short timeout to avoid blocking startup or probes
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis healthcheck failed: %w", err)
	}

	return nil
}