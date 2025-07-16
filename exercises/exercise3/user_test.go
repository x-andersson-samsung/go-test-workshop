package userservice

import (
	"testing"
)

func TestUserService_Create(t *testing.T) {
	s := NewUserService()
	user, err := s.Create("test@example.com", "Test User", 25)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if user.Email != "test@example.com" {
		t.Errorf("expected email %q, got %q", "test@example.com", user.Email)
	}
}

func TestUserService_Get(t *testing.T) {
	s := NewUserService()
	user, _ := s.Create("test@example.com", "Test User", 25)
	found, err := s.Get(user.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if found.ID != user.ID {
		t.Errorf("expected ID %d, got %d", user.ID, found.ID)
	}
}
