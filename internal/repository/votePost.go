package repository

import (
	"context"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
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
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`INSERT INTO likes (user_id, content_type, content_id) VALUES ($1, $2, $3);`)
	if err != nil {
		return fmt.Errorf("repo: like post: prepare - %w", err)
	}

	_, err = stmt.ExecContext(ctx, vote.UserID, vote.ContentType, vote.ContentId)
	if err != nil {
		return fmt.Errorf("repo: like post: exec - %w", err)
	}

	return nil
}

func (r *VotePostRepository) DislikePost(vote model.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`INSERT INTO dislikes (user_id, content_type, content_id) VALUES ($1, $2, $3);`)
	if err != nil {
		return fmt.Errorf("repo: like post: prepare - %w", err)
	}

	_, err = stmt.ExecContext(ctx, vote.UserID, vote.ContentType, vote.ContentId)
	if err != nil {
		return fmt.Errorf("repo: like post: exec - %w", err)
	}

	return nil
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
