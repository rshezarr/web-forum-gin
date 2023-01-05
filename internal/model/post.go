package model

import "time"

type Post struct {
	ID           int       `json:"post_id" db:"id"`
	UserID       int       `json:"post_user_id" db:"user_id"`
	Title        string    `json:"post_title" db:"title"`
	Content      string    `json:"post_content" db:"content"`
	CreationTime time.Time `json:"post_createdAt" db:"creation_time"`
	Likes        int       `json:"post_likes" db:"likes"`
	Dislikes     int       `json:"post_dislikes" db:"dislikes"`
}
