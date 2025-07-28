package service

import (
	"context"
	"errors"
	"fmt"
	"solution4/repository"
)

var (
	DBError = errors.New("db error")
)

// UserRepository defines the interface used during communication with db
//
//go:generate mockgen -destination=mocks/repository.gen.go -package=mocks . UserRepository
type UserRepository interface {
	Create(ctx context.Context, user repository.User) (repository.UserID, error)
	Get(ctx context.Context, email string) (*repository.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email string) (repository.UserID, error) {
	user := repository.User{Name: name, Email: email}
	receivedUser, err := s.repo.Get(ctx, email)
	if err != nil && !errors.Is(err, repository.UserNotFoundError) {
		return -1, errors.Join(DBError, err)
	}

	if receivedUser != nil && receivedUser.Email == email {
		return -1, fmt.Errorf("user with email %q already exists", email)
	}

	createdUserID, err := s.repo.Create(ctx, user)
	if err != nil {
		return -1, errors.Join(DBError, err)
	}
	return createdUserID, nil
}
