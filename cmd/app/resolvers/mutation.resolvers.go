package resolvers

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	graph "go-graphql/graph"
	model "go-graphql/cmd/app/domain"
	model1 "go-graphql/cmd/app/domain/dao"
)

// Notes:
// - Resolvers created so struct that auto defined in generated.go can be used
// - mutationResolver dan queryResolver ada di schema.resolvers.go sama-sama return Resolver

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model1.Post, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	post := &model1.Post{
		ID: fmt.Sprintf("T%d", rand), 
		Title: input.Title,
		Description: input.Description,
	}
	r.Posts = append(r.Posts, post)
	return post, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model1.User, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	user := &model1.User{
		ID: fmt.Sprintf("T%d", rand),
		Name: input.Name, 
		Email: input.Email,
	}
	r.Users = append(r.Users, user)
	return user, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct { *Resolver } 