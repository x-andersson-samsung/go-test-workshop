//go:build gomock && !integration

package service_test

import (
	"context"
	"solution4/repository"
	"testing"

	"solution4/service"
	"solution4/service/mocks"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegisterUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	svc := service.NewUserService(mockRepo)

	ctx := context.Background()
	user := repository.User{Name: "Alice", Email: "alice@example.com"}
	expectedID := repository.UserID(1)

	mockRepo.EXPECT().
		Get(gomock.Any(), user.Email).
		Return(nil, repository.UserNotFoundError)

	mockRepo.EXPECT().
		Create(gomock.Any(), user).
		Return(expectedID, nil)

	userID, err := svc.RegisterUser(ctx, user.Name, user.Email)
	require.NoError(t, err)
	require.Equal(t, expectedID, userID)
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	svc := service.NewUserService(mockRepo)

	ctx := context.Background()
	existingUser := &repository.User{Name: "Alice", Email: "alice@example.com"}

	mockRepo.EXPECT().
		Get(gomock.Any(), "alice@example.com").
		Return(existingUser, nil)

	_, err := svc.RegisterUser(ctx, "Alice 2", "alice@example.com")
	require.Error(t, err)
	require.ErrorContains(t, err, "already exists")
}

func TestRegisterUser_CreationError(t *testing.T) {
	t.Run("empty_name", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().
			Get(gomock.Any(), "a@example.com").
			Return(nil, repository.UserNotFoundError)

		mockRepo.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(repository.UserID(-1), repository.UserCreationError)

		_, err := svc.RegisterUser(context.Background(), "", "a@example.com")
		require.ErrorIs(t, err, repository.UserCreationError)
	})

	t.Run("empty_email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().
			Get(gomock.Any(), "").
			Return(nil, repository.UserNotFoundError)

		mockRepo.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(repository.UserID(-1), repository.UserCreationError)

		_, err := svc.RegisterUser(context.Background(), "Bob", "")
		require.ErrorIs(t, err, repository.UserCreationError)
	})
}
