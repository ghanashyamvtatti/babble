package config

import (
	"pubsubhub/models"
	"time"
)

type ApplicationConfig struct {
	Users         []*models.User
	Posts         []*models.Post
	Subscriptions []*models.Subscription
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Users:         nil,
		Posts:         nil,
		Subscriptions: nil,
	}

	return appConfig
}

func (appConfig *ApplicationConfig) AddUser(fullName string, username string, password string) {
	newUser := &models.User{
		Id:        len(appConfig.Users),
		FullName:  fullName,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	appConfig.Users = append(appConfig.Users, newUser)
}