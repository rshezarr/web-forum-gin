package service

import (
	"forum/internal/repository"
)

type Service struct {
	User
	Post
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUser(repo.User),
		Post: NewPost(repo.Post),
	}
}
