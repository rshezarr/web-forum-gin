package service

import (
	"forum/internal/repository"
)

type Service struct {
	User
	Post
	Comment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUser(repo.User),
		Post:    NewPost(repo.Post),
		Comment: NewComment(repo.Comment),
	}
}
