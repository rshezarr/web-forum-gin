package service

import (
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/internal/repository"
)

var (
	ErrInvalidPostTitle   = errors.New("invalid post title characters")
	ErrInvalidPostContent = errors.New("invalid post content characters")
	ErrPostTitleLen       = errors.New("title length out of range")
	ErrPostContentLen     = errors.New("content length out of range")
	ErrImageSize          = errors.New("image size bigger than 20MB")
	ErrImageType          = errors.New("invalid image type")
)

type Post interface {
	CreatePost(post model.Post) (int, error)
	GetPostByID(postId int) (model.Post, error)
	UpdatePost(newPost model.Post) (int, error)
	DeletePost(postId int) error
	GetAllPosts() ([]model.Post, error)
}

type PostService struct {
	repo repository.Post
}

func NewPost(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func checkPost(post model.Post) error {
	if len(post.Title) > 100 {
		return fmt.Errorf("service: Create Post: check post: %w", ErrPostTitleLen)
	}

	if len(post.Content) > 1500 {
		return fmt.Errorf("service: Create Post: check post: %w", ErrPostContentLen)
	}

	if post.Title == "" {
		return fmt.Errorf("service: Create Post: check post: %w", ErrInvalidPostTitle)
	}

	if post.Content == "" {
		return fmt.Errorf("service: Create Post: check post: %w", ErrInvalidPostContent)
	}

	return nil
}

func (s *PostService) CreatePost(post model.Post) (int, error) {
	if err := checkPost(post); err != nil {
		return 0, err
	}

	id, err := s.repo.CreatePost(post)
	if err != nil {
		return 0, fmt.Errorf("service: create post: %w", err)
	}

	return id, nil
}

func (s *PostService) UpdatePost(newPost model.Post) (int, error) {
	id, err := s.repo.UpdatePost(newPost)
	if err != nil {
		return 0, fmt.Errorf("service: update post: %w", err)
	}

	return id, nil
}

func (s *PostService) DeletePost(newPost model.Post) error {
	return nil
}

func (s *PostService) GetPostByID(postId int) (model.Post, error) {
	post, err := s.repo.GetPostByID(postId)
	if err != nil {
		return model.Post{}, fmt.Errorf("service: get post by id: %w", err)
	}

	return post, nil
}

func (s *PostService) GetAllPosts() ([]model.Post, error) {
	allPosts, err := s.repo.GetAllPosts()
	if err != nil {
		return nil, fmt.Errorf("service: get all posts: %w", err)
	}

	return allPosts, nil
}
