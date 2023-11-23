package model

import "time"

type UserDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserEntity struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type User struct {
	ID             int       `json:"id" db:"id"`
	Email          string    `json:"email" db:"email"`
	Username       string    `json:"username" db:"username"`
	Password       string    `json:"password" db:"password"`
	Posts          int       `json:"posts" db:"posts"`
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"token_expiry"`
}

func UserDtoToEntity(dto *UserDto) *UserEntity {
	return &UserEntity{
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
	}
}
