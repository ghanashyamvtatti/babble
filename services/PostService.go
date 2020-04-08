package services

import (
	"ds-project/DAL"
	"ds-project/config"
	"ds-project/models"
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
	return DAL.GetPosts(appConfig, username)
}

func GetFeedForUsername(appConfig *config.ApplicationConfig, username string) []*models.Post {
	subscriptions := GetSubscriptionsForUsername(appConfig, username)
	var posts []*models.Post

	for _, subscription := range subscriptions {
		posts = append(posts, GetPostsForUser(appConfig, subscription)...)
	}

	return posts
}
