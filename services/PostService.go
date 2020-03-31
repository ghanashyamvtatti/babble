package services

import (
	"ds-project/config"
	"ds-project/models"
	"time"
	"sync"
)

var (
	mutex sync.Mutex
)
func AddPost(appConfig *config.ApplicationConfig, username string, post string) {
	mutex.Lock()
	newPost := &models.Post{
		Post:      post,
		Username:  username,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	appConfig.Posts[username] = append(appConfig.Posts[username], newPost)
	mutex.Unlock()
}

func GetPostsForUser(appConfig *config.ApplicationConfig, username string) []*models.Post {
	return appConfig.Posts[username]
}

func GetFeedForUsername(appConfig *config.ApplicationConfig, username string) []*models.Post {
	subscriptions := GetSubscriptionsForUsername(appConfig, username)
	var posts []*models.Post

	for _, subscription := range subscriptions {
		posts = append(posts, GetPostsForUser(appConfig, subscription)...)
	}

	return posts
}
