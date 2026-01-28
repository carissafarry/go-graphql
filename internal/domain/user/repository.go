package user

type Repository interface {
	FindAll() ([]*User, error)
	Create(name, email string) (*User, error)
}