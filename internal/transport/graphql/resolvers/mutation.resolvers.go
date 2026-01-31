package resolvers

import (
	"context"

	graph "go-graphql/internal/transport/graphql/graph"
	model "go-graphql/internal/transport/graphql/model"
)

// Notes:
// - Resolvers created so struct that auto defined in generated.go can be used
// - mutationResolver dan queryResolver ada di schema.resolvers.go sama-sama return Resolver

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	userID := "1" 							// TODO: hardcoded for demo purposes
	p, err := r.PostUsecase.CreatePost(
		ctx,
		input.Title,
		input.Description,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	u, err := r.UserUsecase.CreateUser(
		ctx,
		input.Name,
		input.Email,
	)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct { *Resolver } 