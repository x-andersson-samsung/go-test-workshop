package service

import (
	"context"
	"errors"
	"fmt"

	"exercise4/repository"
)

//1. Extract Interface
//
//Create an interface for the db connection functionality
//Update service.UserService to use the interface instead of concrete type
//
//2. Create fake implementation and generate mock
//
//Create a fake implementation of the user repository interface
//Generate a mock implementation of the user repository interface using uber-go/mock
//Both implementations should:
//  Allow consumer to create new User
//  Allow consumer to get User by email
//
//3. Write Tests
//
//Test successful user creation
//Test error handling and invalid input handling

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
