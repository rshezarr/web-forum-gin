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

func (s *PostService) CreatePost(post model.Post) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PostService) GetAllPosts() ([]model.Post, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PostService) GetPostByID(postId int) (model.Post, error) {
	panic("not implemented") // TODO: Implement
}
