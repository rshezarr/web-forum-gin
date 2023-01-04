package model

import "time"

type Post struct {
	ID           int `json:"-"`
	UserID       int
	Title        string    `json:"post_title"`
	Content      string    `json:"post_content"`
	CreationTime time.Time `json:"post_createdAt"`
	Category     []string  `json:"post_categories"`
	Likes        int
	Dislikes     int
}
