package postgres

import (
	"database/sql"
	"strconv"

	"go-graphql/internal/domain/post"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

// FindAll implements post.Repository
func (r *PostRepo) FindAll() ([]*post.Post, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, user_id
		FROM posts
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*post.Post
	for rows.Next() {
		p := &post.Post{}
		var id int

		if err := rows.Scan(&id, &p.Title, &p.Description, &p.UserID); err != nil {
			return nil, err
		}
		p.ID = strconv.Itoa(id)
		posts = append(posts, p)
	}
	return posts, nil
}

// FindByUserID implements post.Repository
func (r *PostRepo) FindByUserID(userID string) ([]*post.Post, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, user_id
		FROM posts
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*post.Post
	for rows.Next() {
		p := &post.Post{}
		var id int

		if err := rows.Scan(&id, &p.Title, &p.Description, &p.UserID); err != nil {
			return nil, err
		}
		p.ID = strconv.Itoa(id)
		posts = append(posts, p)
	}
	return posts, nil
}

// Create implements post.Repository
func (r *PostRepo) Create(title, description, userID string) (*post.Post, error) {
	p := &post.Post{
		Title:       title,
		Description: description,
		UserID:      userID,
	}

	var id int
	err := r.db.QueryRow(
		`INSERT INTO posts(title, description, user_id)
		 VALUES($1, $2, $3)
		 RETURNING id`,
		title, description, userID,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	p.ID = strconv.Itoa(id)
	return p, nil
}