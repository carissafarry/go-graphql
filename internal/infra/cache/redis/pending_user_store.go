package redis

import (
	"context"
	"encoding/json"
	"time"

	"go-graphql/internal/domain/user"

	"github.com/redis/go-redis/v9"
)

type PendingUserStore struct {
	client *redis.Client
}

func NewPendingUserStore(client *redis.Client) *PendingUserStore {
	return &PendingUserStore{client: client}
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
	return r.client.SetEx(ctx, key, data, ttl).Err()
}

func (r *PendingUserStore) Find(
	ctx context.Context,
	email string,
) (*user.PendingUser, error) {

	val, err := r.client.Get(ctx, r.key(email)).Result()
	if err != nil {
		return nil, err
	}

	var u user.PendingUser
	if err := json.Unmarshal([]byte(val), &u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PendingUserStore) Delete(
	ctx context.Context,
	email string,
) error {
	return r.client.Del(ctx, r.key(email)).Err()
}

func (r *PendingUserStore) key(email string) string {
	return "pending_user:" + email
}
