package user

import (
	"context"
	"errors"
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetUsers(ctx context.Context) ([]*User, error) {
	return u.repo.FindAll()
}

func (u *Usecase) CreateUser(
	ctx context.Context,
	name string,
	email string,
) (*User, error) {
	// Validation
	if name == "" {
		return nil, errors.New("name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}

	// Orchestration
	return u.repo.Create(name, email)
}