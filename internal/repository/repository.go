package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUser(db),
	}
}
