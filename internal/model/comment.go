package model

type Comment struct {
	ID       int    `json:"comment_id"`
	UserID   int    `json:"comment_user_id"`
	PostID   int    `json:"comment_post_id"`
	Content  string `json:"comment-content"`
	Likes    int    `json:"comment_likes"`
	Dislikes int    `json:"comment_dislikes"`
}
