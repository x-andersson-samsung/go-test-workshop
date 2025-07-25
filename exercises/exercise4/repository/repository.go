package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

var (
	UserNotFoundError = errors.New("user not found")
	UserCreationError = errors.New("failed to create user")
)

type UserID int

type User struct {
	ID    UserID
	Name  string
	Email string
}

type PostgresRepository struct {
	Conn *pgx.Conn
}

func NewPostgresRepository(conn *pgx.Conn) (*PostgresRepository, error) {
	return &PostgresRepository{Conn: conn}, nil
}

func (r *PostgresRepository) Get(ctx context.Context, email string) (*User, error) {
	row := r.Conn.QueryRow(ctx, "SELECT id, name, email FROM users WHERE email=$1", email)
	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, UserNotFoundError // Not found is OK
	}
	return &u, nil
}

func (r *PostgresRepository) Create(ctx context.Context, user User) (UserID, error) {
	if user.Email == "" || user.Name == "" {
		return -1, UserCreationError
	}

	var id int
	err := r.Conn.QueryRow(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return UserID(id), nil
}
