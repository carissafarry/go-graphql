package resolvers

import (
	"errors"

	"go-graphql/internal/domain/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func mapUserError(err error) error {
	switch {
	case errors.Is(err, user.ErrEmailAlreadyRegistered):
		return gqlerror.Errorf("email already registered")

	case errors.Is(err, user.ErrInvalidOTP):
		return gqlerror.Errorf("invalid or expired otp")

	default:
		return gqlerror.Errorf("internal server error")
	}
}
