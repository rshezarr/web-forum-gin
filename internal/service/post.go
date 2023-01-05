package service

import (
	"forum/internal/model"
	"forum/internal/repository"
)

type Post interface {
	CreatePost(post model.Post) (int, error)
	GetAllPosts() ([]model.Post, error)
	GetPostByID(postId int) (model.Post, error)
}

type PostService struct {
	repo repository.Post
}

func NewPost(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}
