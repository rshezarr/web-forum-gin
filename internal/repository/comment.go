package repository

import "github.com/jmoiron/sqlx"

type Comment interface {
}

type CommentRepository struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}
