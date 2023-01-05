package repository

import "github.com/jmoiron/sqlx"

type VotePost interface {
}

type VotePostRepository struct {
	db *sqlx.DB
}

func NewVotePost(db *sqlx.DB) *VotePostRepository {
	return &VotePostRepository{
		db: db,
	}
}
