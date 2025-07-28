//go:build integration

package service_test

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"solution4/repository"
	"testing"

	"solution4/service"
)

const (
	dbName     = "users"
	dbUser     = "user"
	dbPassword = "password"
)

type UserServiceIntegrationSuite struct {
	suite.Suite
	ctx       context.Context
	container *postgres.PostgresContainer
	conn      *pgx.Conn
	snapshot  []byte
}

func (s *UserServiceIntegrationSuite) SetupSuite() {
	s.ctx = context.Background()

	container, err := postgres.Run(s.ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	require.NoError(s.T(), err)
	s.container = container

	_, _, err = container.Exec(s.ctx, []string{"psql", "-U", dbUser, "-d", dbName, "-c", "CREATE TABLE users (id SERIAL PRIMARY KEY,name TEXT NOT NULL,email TEXT NOT NULL UNIQUE);"})
	require.NoError(s.T(), err)

	// Insert baseline user
	_, _, err = container.Exec(s.ctx, []string{"psql", "-U", dbUser, "-d", dbName, "-c", `INSERT INTO users (name, email) VALUES ('Alice', 'alice@example.com')`})
	require.NoError(s.T(), err)

	// Take a snapshot of the current DB
	err = container.Snapshot(s.ctx)
	require.NoError(s.T(), err)
}

func (s *UserServiceIntegrationSuite) SetupTest() {
	err := s.container.Restore(s.ctx)
	require.NoError(s.T(), err)

	dbURL, err := s.container.ConnectionString(s.ctx)
	require.NoError(s.T(), err)
	s.conn, err = pgx.Connect(s.ctx, dbURL)
	require.NoError(s.T(), err)
}

func (s *UserServiceIntegrationSuite) TearDownTest() {
	_ = s.conn.Close(s.ctx)
}

func (s *UserServiceIntegrationSuite) TearDownSuite() {
	_ = s.container.Terminate(s.ctx)
}

func (s *UserServiceIntegrationSuite) TestRegisterUser_Success() {
	repo, err := repository.NewPostgresRepository(s.conn)
	require.NoError(s.T(), err)
	svc := service.NewUserService(repo)

	id, err := svc.RegisterUser(s.ctx, "Alice2", "alice2@example.com")
	require.NoError(s.T(), err)
	require.Equal(s.T(), repository.UserID(2), id)
}

func (s *UserServiceIntegrationSuite) TestRegisterUser_DuplicateEmail() {
	repo, err := repository.NewPostgresRepository(s.conn)
	require.NoError(s.T(), err)
	svc := service.NewUserService(repo)

	_, err = svc.RegisterUser(s.ctx, "Another", "alice@example.com")
	require.Error(s.T(), err)
	require.ErrorContains(s.T(), err, "already exists")
}

func (s *UserServiceIntegrationSuite) TestRegisterUser_CreationError() {
	repo, err := repository.NewPostgresRepository(s.conn)
	require.NoError(s.T(), err)
	svc := service.NewUserService(repo)

	s.T().Run("empty_name", func(t *testing.T) {
		_, err = svc.RegisterUser(s.ctx, "", "x@example.com")
		require.ErrorContains(t, err, repository.UserCreationError.Error())
	})

	s.T().Run("empty_email", func(t *testing.T) {
		_, err = svc.RegisterUser(s.ctx, "Bob", "")
		require.ErrorContains(t, err, repository.UserCreationError.Error())
	})
}

func TestUserServiceIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(UserServiceIntegrationSuite))
}
