package cache

import "forum/internal/model"

type PostCache interface {
	Set(kay string, value *model.Post)
	Get(kay string) *model.Post
}
