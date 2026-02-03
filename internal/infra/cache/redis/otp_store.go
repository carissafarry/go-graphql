package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type OTPStore struct {
	client *redis.Client
}

func NewOTPStore(client *redis.Client) *OTPStore {
	return &OTPStore{client: client}
}

func (r *OTPStore) Save(
	ctx context.Context,
	email string,
	otp string,
	ttl time.Duration,
) error {
	return r.client.SetEx(
		ctx,
		r.key(email),
		otp,
		ttl,
	).Err()
}

func (r *OTPStore) Find(
	ctx context.Context,
	email string,
) (string, error) {
	return r.client.Get(ctx, r.key(email)).Result()
}

func (r *OTPStore) Delete(
	ctx context.Context,
	email string,
) error {
	return r.client.Del(ctx, r.key(email)).Err()
}

func (r *OTPStore) key(email string) string {
	return "otp:" + email
}
