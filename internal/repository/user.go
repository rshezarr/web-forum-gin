package repository

import (
	"forum/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
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
	panic("not implemented") // TODO: Implement
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
