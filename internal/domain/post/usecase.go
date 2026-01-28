package post

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

func (u *Usecase) GetPosts(ctx context.Context) ([]*Post, error) {
	return u.repo.FindAll()
}

func (u *Usecase) CreatePost(
	ctx context.Context,
	title string,
	description string,
	userID string,
) (*Post, error) {
	// Validation
	if title == "" {
		return nil, errors.New("title is required")
	}
	if userID == "" {
		return nil, errors.New("userID is required")
	}

	// Orchestration
	return u.repo.Create(title, description, userID)
}