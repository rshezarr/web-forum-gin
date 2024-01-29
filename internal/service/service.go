package service

import (
	"forum/internal/repository"
	"forum/internal/service/comment_svc"
	"forum/internal/service/post_svc"
	"forum/internal/service/user_svc"
)

type ServiceInitializer interface {
	UserSvcInit() user_svc.Userer
	PostSvcInit() post_svc.Poster
	CommentSvcInit() comment_svc.Commenter
}

type service struct {
	repo repository.RepoInitializer

	user    user_svc.Userer
	post    post_svc.Poster
	comment comment_svc.Commenter
}

func NewService(repo repository.RepoInitializer) ServiceInitializer {
	return &service{
		repo: repo,
	}
}

func (s *service) UserSvcInit() user_svc.Userer {
	if s.user == nil {
		s.user = user_svc.NewUser(s.repo.UserRepoInit())
	}

	return s.user
}

func (s *service) PostSvcInit() post_svc.Poster {
	if s.post == nil {
		s.post = post_svc.NewPost(s.repo.PostRepoInit())
	}

	return s.post
}

func (s *service) CommentSvcInit() comment_svc.Commenter {
	if s.comment == nil {
		s.comment = comment_svc.NewComment(s.repo.CommentRepoInit())
	}

	return s.comment
}
