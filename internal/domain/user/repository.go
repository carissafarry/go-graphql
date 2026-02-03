package user

import "context"

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, u *User) (*User, error)
}