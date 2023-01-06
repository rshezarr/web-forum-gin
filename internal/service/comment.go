package service

import (
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/internal/repository"
)

type Comment interface {
	Create(comment model.Comment) (int, error)
	GetByPostID(postId int) ([]model.Comment, error)
	Update(newComment string, id int) (int, error)
	Delete(id int) error
}

var ErrInvalidComment = errors.New("invalid comment content")

type CommentService struct {
	repo repository.Comment
}

func NewComment(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func checkComment(comment model.Comment) error {
	if comment.Content == "" {
		return fmt.Errorf("service: create comment: check comment: %w", ErrInvalidComment)
	}
	return nil
}

func (s *CommentService) Create(comment model.Comment) (int, error) {
	if err := checkComment(comment); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(comment)
	if err != nil {
		return 0, fmt.Errorf("service: create comment: check comment: %w", err)
	}

	return id, nil
}

func (s *CommentService) GetByPostID(postId int) ([]model.Comment, error) {
	comment, err := s.repo.GetByPostID(postId)
	if err != nil {
		return nil, fmt.Errorf("service: get comment: %w", err)
	}

	return comment, nil
}

func (s *CommentService) Update(newComment string, id int) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (s *CommentService) Delete(id int) error {
	panic("not implemented") // TODO: Implement
}
