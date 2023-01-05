package repository

import "github.com/jmoiron/sqlx"

type VotePost interface {
	LikePost(userId, contentId int, contentType string) error
	DislikePost(userId, contentId int, contentType string) error
	IsLikedPost(userId, contentId int, contentType string) (bool, error)
	IsDislikedPost(userId, contentId int, contentType string) (bool, error)
	RemoveLike(userId, contentId int, contentType string) error
	RemoveDislike(userId, contentId int, contentType string) error
}

type VotePostRepository struct {
	db *sqlx.DB
}

func NewVotePost(db *sqlx.DB) *VotePostRepository {
	return &VotePostRepository{
		db: db,
	}
}

func (r *VotePostRepository) LikePost(userId int, contentId int, contentType string) error {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) DislikePost(userId int, contentId int, contentType string) error {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) IsLikedPost(userId int, contentId int, contentType string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) IsDislikedPost(userId int, contentId int, contentType string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) RemoveLike(userId int, contentId int, contentType string) error {
	panic("not implemented") // TODO: Implement
}

func (r *VotePostRepository) RemoveDislike(userId int, contentId int, contentType string) error {
	panic("not implemented") // TODO: Implement
}
