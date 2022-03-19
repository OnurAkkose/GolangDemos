package main

import (
	"fmt"
	"redisdemo/cache"
)

func main() {
	var postCache cache.PostsCache
	postCache = cache.NewRedisCache("localhost:6379", 1, 0)
	post := entity.Post{
		ID:    "id",
		Title: "title",
	}
	postCache.Set("entity", &post)
	getPost := postCache.Get("entity")
	fmt.Print(*getPost)
}
