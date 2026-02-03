package resolvers

import (
	"context"

	model "go-graphql/internal/transport/graphql/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(
	ctx context.Context,
	input model.NewUser,
) (*model.User, error) {
	u, err := r.UserUsecase.CreateUser(
		ctx,
		input.FullName,
		input.Email,
	)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
	}, nil
}

func (r *mutationResolver) Register(
	ctx context.Context,
	input model.RegisterInput,
) (bool, error) {

	err := r.UserUsecase.Register(
		ctx,
		input.Email,
		input.Password,
	)

	if err != nil {
		return false, mapUserError(err)
	}

	return true, nil
}

func (r *mutationResolver) VerifyOtp(
	ctx context.Context,
	input model.VerifyOTPInput,
) (bool, error) {

	if err := r.UserUsecase.VerifyOTP(
		ctx,
		input.Email,
		input.Otp,
	); err != nil {
		return false, mapUserError(err)
	}

	return true, nil
}