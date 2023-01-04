package model

type Comment struct {
	ID       int
	UserID   int    `json:"comment_user_id"`
	PostID   int    `json:"comment_post_id"`
	Content  string `json:"comment-content"`
	Likes    int
	Dislikes int
}
