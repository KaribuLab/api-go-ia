package main

import (
	"errors"

	"github.com/google/uuid"
)

type UserService interface {
	FindAll() ([]User, error)
	Create(user *User) error
	FindByID(id string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) FindAll() ([]User, error) {
	return s.repo.FindAll()
}

func (s *userService) Create(user *User) error {
	if user.UserName == "" {
		return errors.New("user name is required")
	}

	return s.repo.Create(user)
}

func (s *userService) FindByID(id string) (*User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("invalid user ID")
	}

	user := User{}
	if err := s.repo.FindByID(id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
