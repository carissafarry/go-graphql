package resolvers

import (
	"context"

	graph "go-graphql/graph"
	model1 "go-graphql/cmd/app/domain/dao"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model1.Post, error) {
	return r.Resolver.Posts, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model1.User, error) {
	return r.Resolver.Users, nil
}

// Query returns graph.QueryResolver implementation
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }