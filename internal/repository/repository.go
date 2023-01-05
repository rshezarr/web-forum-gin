package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	User
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUser(db),
		Post: NewPost(db),
	}
}
