package repository

import (
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
)

type VotePost interface {
	LikePost(vote model.Vote) error
	DislikePost(vote model.Vote) error
	IsLikedPost(vote model.Vote) (bool, error)
	IsDislikedPost(vote model.Vote) (bool, error)
	RemoveLike(vote model.Vote) error
	RemoveDislike(vote model.Vote) error
}

type VotePostRepository struct {
	db *sqlx.DB
}

func NewVotePost(db *sqlx.DB) *VotePostRepository {
	return &VotePostRepository{
		db: db,
	}
}

func (r *VotePostRepository) LikePost(vote model.Vote) error {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) DislikePost(vote model.Vote) error {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) IsLikedPost(vote model.Vote) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) IsDislikedPost(vote model.Vote) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) RemoveLike(vote model.Vote) error {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) RemoveDislike(vote model.Vote) error {
	panic("not implemented") // TODO: Implement
}
