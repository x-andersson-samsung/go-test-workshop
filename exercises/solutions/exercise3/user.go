package userservice

import (
	"errors"
	"strings"
	"time"
)

// # Refactor - add sentinel errors
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidEmail = errors.New("invalid password")
	ErrInvalidAge   = errors.New("invalid age")
)

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
	if err := s.validateUser(email, age, false); err != nil {
		return nil, err
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
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) Update(id int, email, name string, age int) error {
	if err := s.validateUser(email, age, true); err != nil {
		return err
	}

	user, err := s.Get(id)
	if err != nil {
		return err
	}

	if email != "" {
		user.Email = email
	}

	if name != "" {
		user.Name = name
	}

	if age != 0 {
		user.Age = age
	}

	return nil
}

func (s *UserService) Delete(id int) error {
	if _, exists := s.users[id]; !exists {
		return ErrUserNotFound
	}
	delete(s.users, id)
	return nil
}

// # Refactor - unify parameter checks
func (s *UserService) validateUser(email string, age int, allowEmpty bool) error {
	if (email == "" && !allowEmpty) || !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}

	if (age == 0 && !allowEmpty) || age < 0 || age > 150 {
		return ErrInvalidAge
	}

	return nil
}
