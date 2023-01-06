package repository

import (
	"context"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Comment interface {
	Create(comment model.Comment) (int, error)
	GetByID(id int) (model.Comment, error)
	GetByUserID(userId int) ([]model.Comment, error)
	GetByPostID(postId int) ([]model.Comment, error)
	Update(userId int) (int, error)
	Delete(userId int) error
}

type CommentRepository struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(comment model.Comment) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`INSERT INTO commentaries (content, user_id, post_id) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("repo: create comment: prepare - %w", err)
	}

	var commentID int
	if err := stmt.GetContext(ctx, &commentID, comment.Content, comment.UserID, comment.PostID); err != nil {
		return 0, fmt.Errorf("repo: create comment: get - %w", err)
	}

	defer stmt.Close()

	return commentID, nil
}

func (r *CommentRepository) GetByID(id int) (model.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, post_id, content FROM commentaries WHERE id = $1;`)
	if err != nil {
		return model.Comment{}, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	var comment model.Comment
	if err := stmt.GetContext(ctx, &comment); err != nil {
		return model.Comment{}, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	return comment, nil
}

func (r *CommentRepository) GetByUserID(userId int) ([]model.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, post_id, content FROM commentaries WHERE user_id = $1;`)
	if err != nil {
		return nil, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	var comments []model.Comment
	if err := stmt.SelectContext(ctx, &comments); err != nil {
		return nil, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	return comments, nil
}

func (r *CommentRepository) GetByPostID(postId int) ([]model.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, post_id, content FROM commentaries WHERE user_id = $1;`)
	if err != nil {
		return nil, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	var comments []model.Comment
	if err := stmt.SelectContext(ctx, &comments); err != nil {
		return nil, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	return comments, nil
}

func (r *CommentRepository) Update(userId int) (int, error) {
	return 0, nil
}

func (r *CommentRepository) Delete(userId int) error {
	return nil
}
