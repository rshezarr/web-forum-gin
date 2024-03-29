package post_repo

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Poster interface {
	Create(post model.Post) (int, error)
	GetByID(id int) (model.Post, error)
	Update(newPost model.Post) (int, error)
	Delete(id int) error
	GetAll() ([]model.Post, error)
}

type postRepository struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) Poster {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post model.Post) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, fmt.Errorf("repo: create post: begin - %w", err)
	}

	//first query
	stmt, err := tx.Preparex(`INSERT INTO posts (title, content, user_id, creation_time) VALUES ($1, $2, $3, $4) RETURNING id;`)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("repo: create post: first query: prepare - %w", err)
	}

	var id int
	if err := stmt.GetContext(ctx, &id, post.Title, post.Content, post.UserID, post.CreationTime); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("repo: create post: first query: get - %w", err)
	}

	defer stmt.Close()

	//second query
	stmt, err = tx.Preparex(`UPDATE users SET posts = posts + 1 WHERE id = $1;`)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("repo: create post: second query: prepare - %w", err)
	}

	_, err = stmt.ExecContext(ctx, post.UserID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("repo: create post: second query: exec - %w", err)
	}

	defer stmt.Close()

	return id, tx.Commit()
}

func (r *postRepository) GetByID(id int) (model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, title, content, creation_time FROM posts WHERE id = $1;`)
	if err != nil {
		return model.Post{}, fmt.Errorf("repo: get post by id: prepare - %w", err)
	}

	var post model.Post
	if err := stmt.GetContext(ctx, &post, id); err != nil {
		return model.Post{}, fmt.Errorf("repo: get post by id: get - %w", err)
	}

	defer stmt.Close()

	return post, nil
}

func (r *postRepository) Update(newPost model.Post) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`UPDATE posts SET title = $1, content = $2 WHERE id = $3 RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("repo: update post: prepare - %w", err)
	}

	var id int
	if err := stmt.GetContext(ctx, &id, newPost.Title, newPost.Content, newPost.ID); err != nil {
		return 0, fmt.Errorf("repo: update post: exec - %w", err)
	}

	defer stmt.Close()

	return id, nil
}

func (r *postRepository) Delete(postId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return fmt.Errorf("repo: delete post: begin - %w", err)
	}

	//first query
	stmt, err := tx.Preparex(`SELECT user_id FROM posts WHERE id = $1`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: delete post: first query: prepare - %w", err)
	}

	var userId int
	if err := stmt.GetContext(ctx, &userId, postId); err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: delete post: first query: get - %w", err)
	}

	defer stmt.Close()

	//second query
	stmt, err = tx.Preparex(`DELETE FROM posts WHERE id = $1;`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: delete post: second query: prepare - %w", err)
	}

	_, err = stmt.ExecContext(ctx, postId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: delete post: second query: exec - %w", err)
	}

	defer stmt.Close()

	//third query
	stmt, err = tx.Preparex(`UPDATE users SET posts = posts - 1 WHERE id = $1`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: delete post: third query: prepare - %w", err)
	}

	_, err = stmt.ExecContext(ctx, userId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: delete post: third query: exec - %w", err)
	}

	defer stmt.Close()

	return tx.Commit()
}

func (r *postRepository) GetAll() ([]model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, title, content, creation_time FROM posts;`)
	if err != nil {
		return nil, fmt.Errorf("repo: get all posts: prepare - %w", err)
	}

	var posts []model.Post
	if err := stmt.SelectContext(ctx, &posts); err != nil {
		return nil, fmt.Errorf("repo: get all posts: select - %w", err)
	}

	defer stmt.Close()

	return posts, nil
}
