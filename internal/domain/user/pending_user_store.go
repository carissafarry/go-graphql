package user

import (
	"context"
	"time"
)

type PendingUser struct {
	Email    string
	Password string
}

type PendingUserStore interface {
	Save(ctx context.Context, user *PendingUser, ttl time.Duration) error
	Find(ctx context.Context, email string) (*PendingUser, error)
	Delete(ctx context.Context, email string) error
}
