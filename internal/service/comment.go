package service

import (
	"forum/internal/model"
	"forum/internal/repository"
)

type Comment interface {
	Create(comment model.Comment) (int, error)
	GetByPostID(postId int) (model.Comment, error)
	Update(newComment string, id int) (int, error)
	Delete(id int) error
}

type CommentService struct {
	repo repository.Comment
}

func NewComment(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}
