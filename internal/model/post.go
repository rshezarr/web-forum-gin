package model

import "time"

type Post struct {
	ID           int `json:"post_id"`
	UserID       int
	Title        string    `json:"post_title"`
	Content      string    `json:"post_content"`
	CreationTime time.Time `json:"post_createdAt"`
	Category     []string  `json:"post_categories"`
	Likes        int       `json:"post_likes"`
	Dislikes     int       `json:"post_dislikes"`
}
