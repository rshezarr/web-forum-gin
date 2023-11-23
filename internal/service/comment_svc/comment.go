package comment_svc

import (
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/internal/repository/comment_repo"
)

type Commenter interface {
	Create(comment model.Comment) (int, error)
	GetByPostID(postId int) ([]model.Comment, error)
	Update(newComment model.Comment, id int) (int, error)
	Delete(id, userId int) error
}

var ErrInvalidComment = errors.New("invalid comment content")

type commentService struct {
	repo comment_repo.Commenter
}

func NewComment(repo comment_repo.Commenter) Commenter {
	return &commentService{
		repo: repo,
	}
}

func checkComment(comment string) error {
	if comment == "" {
		return fmt.Errorf("service: create comment: check comment: %w", ErrInvalidComment)
	}
	return nil
}

func (s *commentService) Create(comment model.Comment) (int, error) {
	if err := checkComment(comment.Content); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(comment)
	if err != nil {
		return 0, fmt.Errorf("service: create comment: check comment: %w", err)
	}

	return id, nil
}

func (s *commentService) GetByPostID(postId int) ([]model.Comment, error) {
	comment, err := s.repo.GetByPostID(postId)
	if err != nil {
		return nil, fmt.Errorf("service: get comment: %w", err)
	}

	return comment, nil
}

func (s *commentService) Update(newComment model.Comment, id int) (int, error) {
	if err := checkComment(newComment.Content); err != nil {
		return 0, err
	}

	commentId, err := s.repo.Update(newComment, id)
	if err != nil {
		return 0, fmt.Errorf("service: update: %w", err)
	}

	return commentId, nil
}

func (s *commentService) Delete(id, userId int) error {
	if err := s.repo.Delete(id, userId); err != nil {
		return fmt.Errorf("service: delete: %w", err)
	}
	return nil
}
