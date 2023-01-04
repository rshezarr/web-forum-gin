package repository

import (
	"context"
	"fmt"
	"forum/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type User interface {
	CreateUser(user model.User) (int, error)
	GetUser(userID int) (model.User, error)
	SaveToken(username, token string, expirationTime time.Time) error
	GetUserByToken(token string) (model.User, error)
	DeleteToken(token string) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user model.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("repo: create user: prepare - %w", err)
	}

	defer stmt.Close()

	var id int
	if err = stmt.GetContext(ctx, &id, user.Email, user.Username, user.Password); err != nil {
		return 0, fmt.Errorf("repo: create user: get - %w", err)
	}

	return id, nil
}

func (r *UserRepository) GetUser(userID int) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepository) SaveToken(username string, token string, expirationTime time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepository) GetUserByToken(token string) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepository) DeleteToken(token string) error {
	panic("not implemented") // TODO: Implement
}
