package services

import (
	"pubsubhub/config"
	"pubsubhub/models"
	"time"
)

func AddPost(appConfig *config.ApplicationConfig, userId int, post string) {
	newPost := &models.Post{
		Id:        len(appConfig.Posts),
		UserId:    userId,
		Post:      post,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	appConfig.Posts = append(appConfig.Posts, newPost)
}

func GetPostsForUserId(appConfig *config.ApplicationConfig, userId int) []*models.Post {
	var posts []*models.Post
	for _, post := range appConfig.Posts {
		if post.UserId == userId {
			posts = append(posts, post)
		}
	}
	return posts
}

func GetFeedForUserId(appConfig *config.ApplicationConfig, userId int) []*models.Post {
	subscriptions := GetSubscriptionsForUserId(appConfig, userId)
	var posts []*models.Post

	for _, subscription := range subscriptions {
		posts = append(posts, GetPostsForUserId(appConfig, subscription.Publisher.Id)...)
	}

	return posts
}