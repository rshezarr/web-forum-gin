package repository

import (
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
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
