package postgres

import (
	"database/sql"
	"strconv"
	user "go-graphql/internal/domain/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db:db}
}

func (r *UserRepo) FindAll() ([]*user.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		var id int
		if err := rows.Scan(&id, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		u.ID = strconv.Itoa(id)
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) Create(name, email string) (*user.User, error) {
	u := &user.User{
		Name:  name,
		Email: email,
	}

	var id int
	err := r.db.QueryRow(
		`INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`,
		name, email,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	u.ID = strconv.Itoa(id)
	return u, err
}