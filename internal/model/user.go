package model

import "time"

type User struct {
	ID       int    `json:"-"`
	Email    string `json:"user_email"`
	Username string `json:"user_username"`
	Password string `json:"user_password"`
	Posts    int

	Token          string
	ExpirationTime time.Time
}
