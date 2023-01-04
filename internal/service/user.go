package service

import (
	"forum/internal/model"
	"forum/internal/repository"
)

type User interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (model.User, error)
	ParseToken(token string) (model.User, error)
	DeleteToken(token string) error
}

type UserService struct {
	repo repository.User
}

func NewUser(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}
