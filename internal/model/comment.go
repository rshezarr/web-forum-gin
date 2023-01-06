package model

type Comment struct {
	ID      int    `json:"comment_id" db:"id"`
	UserID  int    `json:"comment_user_id" db:"user_id"`
	PostID  int    `json:"comment_post_id" db:"post_id"`
	Content string `json:"comment-content" db:"content"`
}
