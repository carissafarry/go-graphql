package user

import "time"

// ===== CORE ENTITY =====
type User struct {
	ID    string
	Email string
	FullName string
	Password string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}


// ===== TEMPORARY REGISTRATION STATE =====
type PendingUser struct {
	Email    string
	Password string
}