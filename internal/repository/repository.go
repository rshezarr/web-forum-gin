package repository

import (
	"forum/internal/repository/comment_repo"
	"forum/internal/repository/post_repo"
	"forum/internal/repository/user_repo"
	"github.com/jmoiron/sqlx"
)

type Repositoryer interface {
	UserRepoInit() user_repo.Userer
	PostRepoInit() post_repo.Poster
	CommentRepoInit() comment_repo.Commenter
}

type Repository struct {
	db *sqlx.DB

	user    user_repo.Userer
	post    post_repo.Poster
	comment comment_repo.Commenter
}

func NewRepository(db *sqlx.DB) Repositoryer {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UserRepoInit() user_repo.Userer {
	if r.user == nil {
		r.user = user_repo.NewUser(r.db)
	}

	return r.user
}

func (r *Repository) PostRepoInit() post_repo.Poster {
	if r.post == nil {
		r.post = post_repo.NewPost(r.db)
	}

	return r.post
}

func (r *Repository) CommentRepoInit() comment_repo.Commenter {
	if r.comment == nil {
		r.comment = comment_repo.NewComment(r.db)
	}

	return r.comment
}
