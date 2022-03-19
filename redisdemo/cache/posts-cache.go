package cache

import "redisdemo/entity"

type PostsCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
