package user

import "errors"

var (
	ErrEmailAlreadyRegistered = errors.New("Email already registered")
	ErrInvalidOTP             = errors.New("invalid or expired otp")
	ErrEmailRequired          = errors.New("Email is required")
	ErrFullNameRequired       = errors.New("Full Name is required")
	ErrInvalidPassword        = errors.New("Invalid Password")
)