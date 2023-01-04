package model

import "time"

type User struct {
	ID       int    `json:"-"`
	Email    string `json:"user-email"`
	Username string `json:"user-username"`
	Password string `json:"user-password"`
	Posts    int    `json:"user-posts"`

	Token          string    `json:"user-token"`
	ExpirationTime time.Time `json:"user-token-expire"`
}
