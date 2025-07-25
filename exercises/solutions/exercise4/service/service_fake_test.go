//go:build !gomock && !integration

package service_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"solution4/repository"
	"solution4/service"
	"testing"
)

type FakeUserRepository struct {
	data   map[string]repository.User
	nextID int
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		data:   make(map[string]repository.User),
		nextID: 1,
	}
}

func (db *FakeUserRepository) Create(_ context.Context, user repository.User) (repository.UserID, error) {
	if user.Email == "" || user.Name == "" {
		return -1, repository.UserCreationError
	}

	user.ID = repository.UserID(db.nextID)
	db.nextID++
	db.data[user.Email] = user
	return user.ID, nil
}

func (db *FakeUserRepository) Get(_ context.Context, email string) (*repository.User, error) {
	if email == "" {
		return nil, repository.UserNotFoundError
	}

	user, ok := db.data[email]
	if !ok {
		return nil, repository.UserNotFoundError
	}
	return &user, nil
}

func TestRegisterUser_Success(t *testing.T) {
	repo := NewFakeUserRepository()
	svc := service.NewUserService(repo)

	userID, err := svc.RegisterUser(context.Background(), "Alice", "alice@example.com")
	require.NoError(t, err, "expected successful user creation")

	require.Equal(t, repository.UserID(1), userID, "expected user ID to be 1")
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
	repo := NewFakeUserRepository()
	svc := service.NewUserService(repo)

	_, _ = svc.RegisterUser(context.Background(), "Alice", "alice@example.com")
	_, err := svc.RegisterUser(context.Background(), "Alice 2", "alice@example.com")
	require.Error(t, err)
	require.ErrorContains(t, err, "already exists")
}

func TestRegisterUser_CreationError(t *testing.T) {
	repo := NewFakeUserRepository()
	svc := service.NewUserService(repo)

	t.Run("empty_name", func(t *testing.T) {
		_, err := svc.RegisterUser(context.Background(), "", "a@example.com")
		require.ErrorIs(t, err, repository.UserCreationError)
	})

	t.Run("empty_email", func(t *testing.T) {
		_, err := svc.RegisterUser(context.Background(), "Bob", "")
		require.ErrorIs(t, err, repository.UserCreationError)
	})
}
