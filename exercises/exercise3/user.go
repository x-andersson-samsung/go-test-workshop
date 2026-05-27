package userservice

import (
	"errors"
	"strings"
	"time"
)

//1. Run the coverage analysis:
//
//go test -coverprofile=coverage.out ./...
//go tool cover -func=coverage.out
//go tool cover -html=coverage.out
//
//2. Identify untested code paths:
//
//Which methods need more test coverage?
//What edge cases are missing?
//Are there error conditions not being tested?
//
//3. Improve the test coverage:
//
//Add tests for Update and Delete methods
//Add tests for error conditions
//Test edge cases (empty values, invalid inputs)
//Use table-driven tests where appropriate

type User struct {
	ID        int
	Email     string
	Name      string
	Age       int
	CreatedAt time.Time
}

type UserService struct {
	users  map[int]*User
	lastID int
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

func (s *UserService) Create(email, name string, age int) (*User, error) {
	if email == "" || !strings.Contains(email, "@") {
		return nil, errors.New("invalid email")
	}

	if age < 0 || age > 150 {
		return nil, errors.New("invalid age")
	}

	s.lastID++
	user := &User{
		ID:        s.lastID,
		Email:     email,
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
	}
	s.users[user.ID] = user
	return user, nil
}

func (s *UserService) Get(id int) (*User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) Update(id int, email, name string, age int) error {
	user, err := s.Get(id)
	if err != nil {
		return err
	}

	if email != "" {
		if !strings.Contains(email, "@") {
			return errors.New("invalid email")
		}
		user.Email = email
	}

	if name != "" {
		user.Name = name
	}

	if age != 0 {
		if age < 0 || age > 150 {
			return errors.New("invalid age")
		}
		user.Age = age
	}

	return nil
}

func (s *UserService) Delete(id int) error {
	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
