package model

type Vote struct {
	UserID      int    `json:"vote_user_id"`
	ContentId   int    `json:"vote_content_id"`
	ContentType string `json:"vote_content_type"`
}
