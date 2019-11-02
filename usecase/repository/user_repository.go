package repository

import (
	"../../domain"
)

type UserRepository interface {
	Store(user domain.User) error
	FindByAuthToken(authToken string) (*domain.User, error)
	FindByUserID(userID string) (*domain.User, error)
	UpdateByUserID(userID string, name string) error
}