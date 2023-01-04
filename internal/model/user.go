package model

import "time"

type User struct {
	ID             int       `json:"user_id"`
	Email          string    `json:"user_email"`
	Username       string    `json:"user_username"`
	Password       string    `json:"user_password"`
	Posts          int       `json:"user_posts"`
	Token          string    `json:"user_token"`
	ExpirationTime time.Time `json:"user_expiry"`
}
