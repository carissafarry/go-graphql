package resolvers

import (
	"go-graphql/internal/domain/post"
	"go-graphql/internal/domain/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

// Notes: Resolver will be used in {*}.resolvers.go

type Resolver struct{
	// UserRepo user.Repository
	// PostRepo post.Repository

	UserUsecase *user.Usecase
	PostUsecase *post.Usecase
}
