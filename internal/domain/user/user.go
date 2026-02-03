package user

import "time"

type User struct {
	ID    string
	Email string
	FullName string
	Password string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}