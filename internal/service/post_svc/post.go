package post_svc

import (
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/internal/repository/post_repo"
)

var (
	ErrInvalidPostTitle   = errors.New("invalid post title characters")
	ErrInvalidPostContent = errors.New("invalid post content characters")
	ErrPostTitleLen       = errors.New("title length out of range")
	ErrPostContentLen     = errors.New("content length out of range")
)

type Poster interface {
	Create(post model.Post) (int, error)
	GetByID(postId int) (model.Post, error)
	Update(newPost model.Post) (int, error)
	Delete(postId int) error
	GetAll() ([]model.Post, error)
}

type postService struct {
	repo post_repo.Poster
}

func NewPost(repo post_repo.Poster) Poster {
	return &postService{
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

func (s *postService) Create(post model.Post) (int, error) {
	if err := checkPost(post); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(post)
	if err != nil {
		return 0, fmt.Errorf("service: create post: %w", err)
	}

	return id, nil
}

func (s *postService) GetByID(postId int) (model.Post, error) {
	post, err := s.repo.GetByID(postId)
	if err != nil {
		return model.Post{}, fmt.Errorf("service: get post by id: %w", err)
	}

	return post, nil
}

func (s *postService) Update(newPost model.Post) (int, error) {
	id, err := s.repo.Update(newPost)
	if err != nil {
		return 0, fmt.Errorf("service: update post: %w", err)
	}

	return id, nil
}

func (s *postService) Delete(postId int) error {
	if err := s.repo.Delete(postId); err != nil {
		return fmt.Errorf("service: delete post: %w", err)
	}

	return nil
}

func (s *postService) GetAll() ([]model.Post, error) {
	allPosts, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("service: get all posts: %w", err)
	}

	return allPosts, nil
}
