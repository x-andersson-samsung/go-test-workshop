package userservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTestUser() *User {
	return &User{
		ID:        0,
		Email:     "user@test.com",
		Name:      "Test User",
		Age:       25,
		CreatedAt: time.Now(),
	}
}

func setupTestUserService(users ...*User) *UserService {
	s := NewUserService()
	for _, user := range users {
		// Note: Not the best way of initializing users.
		// It leans a bit too much into testing implementation.
		// We could use `UserService.Create` instead.
		// However, this would mean we are testing 2 functionalities in
		// `Get` / `Delete` / `Update` tests.
		// Switching implementation to use an interface instead of a map
		// would allow us to decouple the service from storage and use mocks
		// for seeding the tests with users for testing.
		// That concept is introduced in Part 4.
		s.users[user.ID] = user
	}
	return s
}

func assertUserInService(t *testing.T, s *UserService, expected *User) {
	// Note: Similar to above this one is going a bit too deep into testing
	// implementation.
	t.Helper()
	got, ok := s.users[expected.ID]

	require.True(t, ok, "user does not exist")

	// Override runtime values. Do it only if you don't care about specific values here
	got.CreatedAt = expected.CreatedAt

	require.Equal(t, expected, got)
}

func TestUserService_Create(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		usr := createTestUser()

		s := setupTestUserService()
		got, err := s.Create(usr.Email, usr.Name, usr.Age)

		require.NoError(t, err)

		// Check if return is correct
		require.Equal(t, usr.Email, got.Email)
		require.Equal(t, usr.Name, got.Name)
		require.Equal(t, usr.Age, got.Age)
		// Not checking createdAt - it is a runtime value

		// Check if the user was saved
		usr.ID = got.ID
		assertUserInService(t, s, usr)
	})
	t.Run("error", func(t *testing.T) {
		cases := map[string]struct {
			email   string
			age     int
			wantErr error
		}{
			"invalid_email": {"test", 25, ErrInvalidEmail},
			"invalid_age":   {"user@test.com", 0, ErrInvalidAge},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				usr := createTestUser()
				usr.Email = tc.email
				usr.Age = tc.age

				s := setupTestUserService()
				_, err := s.Create(usr.Email, usr.Name, usr.Age)
				require.ErrorIs(t, err, tc.wantErr)
			})
		}
	})
}

func TestUserService_Get(t *testing.T) {
	// common setup - our tests are not going to modify state
	usr := createTestUser()
	srv := setupTestUserService(usr)

	t.Run("ok", func(t *testing.T) {
		got, err := srv.Get(usr.ID)

		require.NoError(t, err)

		// Check if return is correct
		require.Equal(t, usr.Email, got.Email)
		require.Equal(t, usr.Name, got.Name)
		require.Equal(t, usr.Age, got.Age)
	})
	t.Run("error_not_found", func(t *testing.T) {
		_, err := srv.Get(usr.ID + 1)

		require.ErrorIs(t, err, ErrUserNotFound)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		usr := createTestUser()
		srv := setupTestUserService(usr)
		err := srv.Delete(usr.ID)

		require.NoError(t, err)
	})
	t.Run("error_not_found", func(t *testing.T) {
		usr := createTestUser()
		srv := setupTestUserService(usr)
		err := srv.Delete(usr.ID + 1)

		require.ErrorIs(t, err, ErrUserNotFound)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		oldUsr := createTestUser()
		newUsr := createTestUser()
		newUsr.Name = "New User"
		newUsr.Email = "new@test.com"
		newUsr.Age = 30

		srv := setupTestUserService(oldUsr)
		err := srv.Update(newUsr.ID, newUsr.Email, newUsr.Name, newUsr.Age)
		require.NoError(t, err)
		assertUserInService(t, srv, newUsr)
	})
	t.Run("error", func(t *testing.T) {
		cases := map[string]struct {
			id      int
			email   string
			age     int
			wantErr error
		}{
			"invalid_email": {0, "test", 25, ErrInvalidEmail},
			"invalid_age":   {0, "user@test.com", -1, ErrInvalidAge},
			"not_found":     {0, "test@test.com", 25, ErrUserNotFound},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				s := setupTestUserService()
				err := s.Update(tc.id, tc.email, "Name", tc.age)
				require.ErrorIs(t, err, tc.wantErr)
			})
		}
	})
}
