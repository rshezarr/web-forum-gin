package model

import "time"

type Post struct {
	ID           int       `json:"post_id"`
	UserID       int       `json:"post_user_id"`
	Title        string    `json:"post_title"`
	Content      string    `json:"post_content"`
	CreationTime time.Time `json:"post_createdAt"`
	Likes        int       `json:"post_likes"`
	Dislikes     int       `json:"post_dislikes"`
}
