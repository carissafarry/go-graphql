package user

import "errors"

var (
	ErrEmailAlreadyRegistered = errors.New("Email already registered")
	ErrInvalidOTP             = errors.New("invalid or expired otp")
)