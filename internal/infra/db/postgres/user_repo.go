package postgres

import (
	"context"
	"database/sql"
	"strconv"
	"time"
	user "go-graphql/internal/domain/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db:db}
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM users WHERE email = $1",
		email,
	)
	
	u := &user.User{}
	var id int
	err := row.Scan(&id, &u.FullName, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // IMPORTANT: not found ≠ error
		}

		return nil, err
	}

	u.ID = strconv.Itoa(id)
	return u, nil
}

func (r *UserRepo) FindAll(ctx context.Context) ([]*user.User, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, full_name, email FROM users`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		var id int

		if err := rows.Scan(&id, &u.FullName, &u.Email); err != nil {
			return nil, err
		}

		u.ID = strconv.Itoa(id)
		users = append(users, u)
	}
	
	return users, nil
}

func (r *UserRepo) Create(ctx context.Context, user *user.User) (*user.User, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}
	now := time.Now().In(loc).Format("15:04:05")

	var id int
	err = r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO users (
			email, full_name, password, role, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4,
			(CURRENT_DATE + $5::time),
			(CURRENT_DATE + $5::time)
		)
		RETURNING id, created_at, updated_at;
		`,
		user.Email,
		user.FullName,
		user.Password,
		user.Role,
		now,
	).Scan(
		&id,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	user.ID = strconv.Itoa(id)
	return user, nil
}