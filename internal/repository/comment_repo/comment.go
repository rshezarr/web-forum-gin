package comment_repo

import (
	"context"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Commenter interface {
	Create(comment model.Comment) (int, error)
	GetByID(id int) (model.Comment, error)
	GetByUserID(userId int) ([]model.Comment, error)
	GetByPostID(postId int) ([]model.Comment, error)
	Update(newComment model.Comment, id int) (int, error)
	Delete(id, userId int) error
}

type commentRepository struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) Commenter {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment model.Comment) (int, error) {
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

func (r *commentRepository) GetByID(id int) (model.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, post_id, content FROM commentaries WHERE id = $1;`)
	if err != nil {
		return model.Comment{}, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	var comment model.Comment
	if err := stmt.GetContext(ctx, &comment); err != nil {
		return model.Comment{}, fmt.Errorf("repo: get comment by id: get - %w", err)
	}

	defer stmt.Close()

	return comment, nil
}

func (r *commentRepository) GetByUserID(userId int) ([]model.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, post_id, content FROM commentaries WHERE user_id = $1;`)
	if err != nil {
		return nil, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	var comments []model.Comment
	if err := stmt.SelectContext(ctx, &comments); err != nil {
		return nil, fmt.Errorf("repo: get comment by id: select - %w", err)
	}

	defer stmt.Close()

	return comments, nil
}

func (r *commentRepository) GetByPostID(postId int) ([]model.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, post_id, content FROM commentaries WHERE post_id = $1;`)
	if err != nil {
		return nil, fmt.Errorf("repo: get comment by id: prepare - %w", err)
	}

	var comments []model.Comment
	if err := stmt.SelectContext(ctx, &comments, postId); err != nil {
		return nil, fmt.Errorf("repo: get comment by id: select - %w", err)
	}

	defer stmt.Close()

	return comments, nil
}

func (r *commentRepository) Update(newComment model.Comment, id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`UPDATE commentaries SET content = $1 WHERE id = $2 AND user_id = $3 RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("repo: update comment: prepare - %w", err)
	}

	var commentId int
	if err := stmt.GetContext(ctx, &commentId, newComment.Content, id, newComment.UserID); err != nil {
		return 0, fmt.Errorf("repo: update comment: get - %w", err)
	}

	defer stmt.Close()

	return commentId, nil
}

func (r *commentRepository) Delete(id, userId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`DELETE FROM commentaries WHERE id = $1 AND user_id = $2;`)
	if err != nil {
		return fmt.Errorf("repo: update comment: prepare - %w", err)
	}

	_, err = stmt.ExecContext(ctx, id, userId)
	if err != nil {
		return fmt.Errorf("repo: update comment: exec - %w", err)
	}

	defer stmt.Close()

	return nil
}
