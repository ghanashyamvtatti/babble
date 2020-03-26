package config

import (
	"ds-project/models"
)

type ApplicationConfig struct {
	Users         map[string]models.User
	Tokens        map[string]string
	Posts         map[string][]*models.Post
	Subscriptions map[string][]string
}

func NewAppConfig() *ApplicationConfig {
	appConfig := &ApplicationConfig{
		Users:         map[string]models.User{},
		Tokens:        map[string]string{},
		Posts:         map[string][]*models.Post{},
		Subscriptions: map[string][]string{},
	}

	return appConfig
}
