package services

import (
	"ds-project/DAL"
	"ds-project/config"
	"ds-project/models"
	"sort"
	"sync"
)

var (
	mutex sync.Mutex
)

func AddPost(appConfig *config.ApplicationConfig, username string, post string) {
	mutex.Lock()
	DAL.AddPost(appConfig, username, post)
	mutex.Unlock()
}

func GetPostsForUser(appConfig *config.ApplicationConfig, username string) []*models.Post {
	posts := DAL.GetPosts(appConfig, username)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	return posts
}

func GetFeedForUsername(appConfig *config.ApplicationConfig, username string) []*models.Post {
	subscriptions := GetSubscriptionsForUsername(appConfig, username)
	var posts []*models.Post

	for _, subscription := range subscriptions {
		posts = append(posts, GetPostsForUser(appConfig, subscription)...)
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	return posts
}
