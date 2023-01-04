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

func (s *UserService) CreateUser(user model.User) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserService) GenerateToken(username string, password string) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserService) ParseToken(token string) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserService) DeleteToken(token string) error {
	panic("not implemented") // TODO: Implement
}
