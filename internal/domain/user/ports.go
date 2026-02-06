package user

import (
	"context"
	"time"
)

// ===== Validation =====
type Validator interface {
	ValidateCreateUser(fullName, email string) error
	ValidateRegister(email, password string) error
	ValidateVerifyOTP(email, otp string) error
}


// ===== Persistence =====
type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, u *User) (*User, error)
}


// ===== Cache / Temporary Storage =====
type PendingUserStore interface {
	Save(ctx context.Context, user *PendingUser, ttl time.Duration) error
	Find(ctx context.Context, email string) (*PendingUser, error)
	Delete(ctx context.Context, email string) error
}

type OTPStore interface {
	Save(ctx context.Context, email string, otp string, ttl time.Duration) error
	Find(ctx context.Context, email string) (string, error)
	Delete(ctx context.Context, email string) error
}


// ===== Utility =====
type OTPGenerator interface {
	GenerateOTP() (string, error)
}