package redis

import "errors"

var (
	ErrOTPNotFound         = errors.New("OTP not found")
	ErrPendingUserNotFound = errors.New("Pending User not found")
)
