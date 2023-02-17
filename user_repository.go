package main

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]User, error)
	Create(user *User) error
	FindByID(id string, user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByID(id string, user *User) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New("invalid user ID")
	}
	return r.db.Where("id = ?", id).First(user).Error
}
