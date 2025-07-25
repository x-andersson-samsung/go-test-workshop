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

type UserService struct {
	repo repository.PostgresRepository
}

func NewUserService(repo repository.PostgresRepository) *UserService {
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
