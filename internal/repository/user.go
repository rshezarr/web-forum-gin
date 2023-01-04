package repository

import (
	"context"
	"fmt"
	"forum/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type User interface {
	CreateUser(user model.User) (int, error)
	GetUserByID(userID int) (model.User, error)
	GetUser(email, hashedPassword string) (model.User, error)
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

func (r *UserRepository) GetUserByID(userID int) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, email, username, password FROM users WHERE id = $1;`)
	if err != nil {
		return model.User{}, fmt.Errorf("repo: get user: prepare - %w", err)
	}

	var user model.User
	if err := stmt.GetContext(ctx, &user, userID); err != nil {
		return model.User{}, fmt.Errorf("repo: ger user: get - %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUser(email, hashedPassword string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("database.ctxTimeout"))
	defer cancel()

	stmt, err := r.db.Preparex(`SELECT id, email, username, password FROM users WHERE email = $1 AND password = $2;`)
	if err != nil {
		return model.User{}, fmt.Errorf("repo: get user: prepare - %w", err)
	}

	var user model.User
	if err := stmt.GetContext(ctx, &user, email, hashedPassword); err != nil {
		return model.User{}, fmt.Errorf("repo: ger user: get - %w", err)
	}

	return user, nil
}
