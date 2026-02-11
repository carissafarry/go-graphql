package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"go-graphql/internal/domain/user"
)

type PendingUserStore struct {
	cache Cache
}

func NewPendingUserStore(cache Cache) *PendingUserStore {
	return &PendingUserStore{cache: cache}
}

func (r *PendingUserStore) Save(
	ctx context.Context,
	u *user.PendingUser,
	ttl time.Duration,
) error {

	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	key := r.key(u.Email)
	return r.cache.Set(ctx, key, data, ttl)
}

func (r *PendingUserStore) Find(
	ctx context.Context,
	email string,
) (*user.PendingUser, error) {

	data, err := r.cache.Get(ctx, r.key(email))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrPendingUserNotFound
		}
		return nil, err
	}

	var u user.PendingUser
	if err := json.Unmarshal(data, &u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PendingUserStore) Delete(
	ctx context.Context,
	email string,
) error {
	return r.cache.Delete(ctx, r.key(email))
}

func (r *PendingUserStore) key(email string) string {
	return "pending_user:" + email
}
