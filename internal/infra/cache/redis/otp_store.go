package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type OTPStore struct {
	cache Cache
}

func NewOTPStore(cache Cache) *OTPStore {
	return &OTPStore{cache: cache}
}

func (r *OTPStore) Save(
	ctx context.Context,
	email string,
	otp string,
	ttl time.Duration,
) error {
	return r.cache.Set(
		ctx,
		r.key(email),
		[]byte(otp),
		ttl,
	)
}

func (r *OTPStore) Find(
	ctx context.Context,
	email string,
) (string, error) {
	data, err := r.cache.Get(ctx, r.key(email))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrOTPNotFound
		}
		return "", err
	}
	return string(data), nil
}

func (r *OTPStore) Delete(
	ctx context.Context,
	email string,
) error {
	return r.cache.Delete(ctx, r.key(email))
}

func (r *OTPStore) key(email string) string {
	return "otp:" + email
}
