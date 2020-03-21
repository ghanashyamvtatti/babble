package config

import (
	"pubsubhub/models"
	"time"
)

type ApplicationConfig struct {
	users         []*models.User
	posts         []*models.Post
	subscriptions []*models.Subscription
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		users:         nil,
		posts:         nil,
		subscriptions: nil,
	}

	return appConfig
}

// Using this as temporary DAL I figure out a more elegant way
func (appConfig *ApplicationConfig) AddUser(fullName string, username string, password string) {
	newUser := &models.User{
		Id:        len(appConfig.users),
		FullName:  fullName,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	appConfig.users = append(appConfig.users, newUser)
}

func (appConfig *ApplicationConfig) GetUserById(userId int) *models.User {
	for _, user := range appConfig.users {
		if user.Id == userId {
			return user
		}
	}
	return nil
}

func (appConfig *ApplicationConfig) AuthenticateUser(username string, password string) *models.User {
	for _, user := range appConfig.users {
		if user.Username == username && user.Password == password {
			return user
		}
	}
	return nil
}

func (appConfig *ApplicationConfig) AddPost(userId int, post string) {
	newPost := &models.Post{
		Id:        len(appConfig.posts),
		UserId:    userId,
		Post:      post,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	appConfig.posts = append(appConfig.posts, newPost)
}

func (appConfig *ApplicationConfig) GetPostsForUserId(userId int) []*models.Post {
	var posts []*models.Post
	for _, post := range appConfig.posts {
		if post.UserId == userId {
			posts = append(posts, post)
		}
	}
	return posts
}

func (appConfig *ApplicationConfig) Subscribe(subscriberId int, publisherId int) {
	newSubscription := &models.Subscription{
		Id:         len(appConfig.subscriptions),
		Subscriber: appConfig.GetUserById(subscriberId),
		Publisher:  appConfig.GetUserById(publisherId),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	appConfig.subscriptions = append(appConfig.subscriptions, newSubscription)
}

func (appConfig *ApplicationConfig) Unsubscribe(subscriberId int, publisherId int) {
	for i, subscription := range appConfig.subscriptions {
		if subscription.Subscriber.Id == subscriberId && subscription.Publisher.Id == publisherId {
			appConfig.subscriptions = append(appConfig.subscriptions[:i], appConfig.subscriptions[i+1:]...)
			return
		}
	}
}

func (appConfig *ApplicationConfig) GetSubscriptionsForUserId(userId int) []*models.Subscription {
	var subscriptions []*models.Subscription
	for _, subscription := range appConfig.subscriptions {
		if subscription.Subscriber.Id == userId {
			subscriptions = append(subscriptions, subscription)
		}
	}
	return subscriptions
}

func (appConfig *ApplicationConfig) GetFeedForUserId(userId int) []*models.Post {
	subscriptions := appConfig.GetSubscriptionsForUserId(userId)
	var posts []*models.Post

	for _, subscription := range subscriptions {
		posts = append(posts, appConfig.GetPostsForUserId(subscription.Publisher.Id)...)
	}

	return posts
}
