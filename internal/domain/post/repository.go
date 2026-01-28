package post

type Repository interface {
	FindAll() ([]*Post, error)
	FindByUserID(userID string) ([]*Post, error)
	Create(title, description, userID string) (*Post, error)
}