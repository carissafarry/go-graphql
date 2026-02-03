package user

import (
	"context"
	"time"
)

type OTPStore interface {
	Save(ctx context.Context, email string, otp string, ttl time.Duration) error
	Find(ctx context.Context, email string) (string, error)
	Delete(ctx context.Context, email string) error
}
