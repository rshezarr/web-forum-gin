package repository

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Post interface {
	CreatePost(post model.Post) (int, error)
	GetPostByID(postId int) (model.Post, error)
	UpdatePost(newPost model.Post) (int, error)
	DeletePost(postId int) (int, error)
	GetAllPosts() ([]model.Post, error)
}

type PostRepository struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(post model.Post) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return 0, fmt.Errorf("repo: create post: begin - %w", err)
	}

	//first query
	stmt, err := tx.Preparex(`INSERT INTO posts (title, content, user_id) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("repo: create post: first query: prepare - %w", err)
	}

	var postID int
	if err := stmt.GetContext(ctx, &postID, post.Title, post.Content, post.UserID); err != nil {
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

	return postID, fmt.Errorf("repo: create post: commit - %w", tx.Commit())
}

func (r *PostRepository) GetPostByID(postId int) (model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, title, content, creation_time, likes, dislikes FROM posts WHERE id = $1;`)
	if err != nil {
		return model.Post{}, fmt.Errorf("repo: get post by id: prepare - %w", err)
	}

	var post model.Post
	if err := stmt.GetContext(ctx, &post, postId); err != nil {
		return model.Post{}, fmt.Errorf("repo: get post by id: get - %w", err)
	}

	return post, nil
}

func (r *PostRepository) UpdatePost(newPost model.Post) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`UPDATE posts SET title = $1, content = $2 WHERE id = $3 RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("repo: update post: prepare - %w", err)
	}

	var userId int
	if err := stmt.GetContext(ctx, &userId, newPost.Title, newPost.Content, newPost.ID); err != nil {
		return 0, fmt.Errorf("repo: update post: exec - %w", err)
	}

	return userId, nil
}

func (r *PostRepository) DeletePost(postId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
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

	return tx.Commit()
}

func (r *PostRepository) GetAllPosts() ([]model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, user_id, title, content, creation_time, likes, dislikes FROM posts;`)
	if err != nil {
		return nil, fmt.Errorf("repo: get all posts: prepare - %w", err)
	}

	var posts []model.Post
	if err := stmt.SelectContext(ctx, &posts); err != nil {
		return nil, fmt.Errorf("repo: get all posts: select - %w", err)
	}

	return posts, nil
}
