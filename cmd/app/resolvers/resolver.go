package resolvers

import model "go-graphql/cmd/app/domain/dao"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

// Notes: Resolver will be used in schema.resolvers.go

type Resolver struct{
	Users []*model.User
	Posts []*model.Post
}
