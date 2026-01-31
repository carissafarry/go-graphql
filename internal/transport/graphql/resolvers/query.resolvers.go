package resolvers

import (
	"context"

	graph "go-graphql/internal/transport/graphql/graph"
	model "go-graphql/internal/transport/graphql/model"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	// return r.PostRepo.FindAll()

	posts, err := r.PostUsecase.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Post
	for _, p := range posts {
		result = append(result, &model.Post{
			ID: p.ID,
			Title: p.Title,
			Description: p.Description,
		})
	}

	return result, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// return r.UserRepo.FindAll()

	users, err := r.UserUsecase.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.User
	for _, u := range users {
		result = append(result, &model.User{
			ID: u.ID, 
			Name: u.Name, 
			Email: u.Email,
		})
	}

	return result, nil
}

// Query returns graph.QueryResolver implementation
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }