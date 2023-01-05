package model

import "time"

type User struct {
	ID             int       `json:"user_id" db:"id"`
	Email          string    `json:"user_email" db:"email"`
	Username       string    `json:"user_username" db:"username"`
	Password       string    `json:"user_password" db:"password"`
	Posts          int       `json:"user_posts" db:"posts"`
	Token          string    `json:"user_token"`
	ExpirationTime time.Time `json:"user_expiry"`
}
