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
	GetAllPosts() ([]model.Post, error)
	GetPostByID(postId int) (model.Post, error)
	GetPostsByCategory(category string) ([]model.Post, error)
	GetCategoriesByPostID(postId int) ([]string, error)
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
		return 0, fmt.Errorf("repo: create post: second query: prepare - %w", err)
	}

	//third query
	for _, category := range post.Category {
		stmt, err := tx.Preparex(`INSERT INTO post_category (post_id, category) VALUES ($1, $2);`)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("repo: create post: third query: prepare - %w", err)
		}

		_, err = stmt.ExecContext(ctx, category)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("repo: create post: third query: prepare - %w", err)
		}
	}

	return postID, nil
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

func (r *PostRepository) GetPostByID(postId int) (model.Post, error) {
	panic("not implemented") // TODO: Implement
}

func (r *PostRepository) GetPostsByCategory(category string) ([]model.Post, error) {
	panic("not implemented") // TODO: Implement
}

func (r *PostRepository) GetCategoriesByPostID(postId int) ([]string, error) {
	panic("not implemented") // TODO: Implement
}
