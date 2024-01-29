package repository

import (
	"forum/internal/repository/comment_repo"
	"forum/internal/repository/post_repo"
	"forum/internal/repository/user_repo"

	"github.com/jmoiron/sqlx"
)

type RepoInitializer interface {
	UserRepoInit() user_repo.Userer
	PostRepoInit() post_repo.Poster
	CommentRepoInit() comment_repo.Commenter
}

type repository struct {
	db *sqlx.DB

	user    user_repo.Userer
	post    post_repo.Poster
	comment comment_repo.Commenter
}

func NewRepository(db *sqlx.DB) RepoInitializer {
	return &repository{
		db: db,
	}
}

func (r *repository) UserRepoInit() user_repo.Userer {
	if r.user == nil {
		r.user = user_repo.NewUser(r.db)
	}

	return r.user
}

func (r *repository) PostRepoInit() post_repo.Poster {
	if r.post == nil {
		r.post = post_repo.NewPost(r.db)
	}

	return r.post
}

func (r *repository) CommentRepoInit() comment_repo.Commenter {
	if r.comment == nil {
		r.comment = comment_repo.NewComment(r.db)
	}

	return r.comment
}
