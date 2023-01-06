package repository

import (
	"context"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
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

func (r *CommentRepository) CreateCommentary(comment model.Comment) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`INSERT INTO commentaries (content, user_id, post_id) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("repo: create post: first query: prepare - %w", err)
	}

	var commentID int
	if err := stmt.GetContext(ctx, &commentID, comment.Content, comment.UserID, comment.PostID); err != nil {
		return 0, fmt.Errorf("repo: create post: first query: get - %w", err)
	}

	defer stmt.Close()

	return commentID, nil
}

func (r *CommentRepository) GetCommentaryByID(id int) (model.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (r *CommentRepository) GetCommentariesByPostID(postId int) ([]model.Comment, error) {
	panic("not implemented") // TODO: Implement
}
