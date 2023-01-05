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

func (r *PostRepository) CreatePost(post model.Post) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (r *PostRepository) GetAllPosts() ([]model.Post, error) {
	panic("not implemented") // TODO: Implement
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
