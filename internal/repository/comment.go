package repository

import (
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
)

type Comment interface {
	CreateCommentary(comment model.Comment) error
	GetCommentaryByID(id int) (model.Comment, error)
	GetCommentariesByPostID(postId int) ([]model.Comment, error)
}

type CommentRepository struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) CreateCommentary(comment model.Comment) error {
	panic("not implemented") // TODO: Implement
}

func (r *CommentRepository) GetCommentaryByID(id int) (model.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (r *CommentRepository) GetCommentariesByPostID(postId int) ([]model.Comment, error) {
	panic("not implemented") // TODO: Implement
}
