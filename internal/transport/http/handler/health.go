package handler

import (
	"context"
	"net/http"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"go-graphql/internal/infra/cache/redis"
)

type HealthHandler struct {
	redisClient *goredis.Client
}

func NewHealthHandler(redisClient *goredis.Client) *HealthHandler {
	return &HealthHandler{redisClient: redisClient}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := redis.Healthcheck(ctx, h.redisClient); err != nil {
		http.Error(w, "Redis unhealthy", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Redis is healthy"))
}