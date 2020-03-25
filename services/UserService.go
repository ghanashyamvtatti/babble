package services

import (
	"pubsubhub/config"
	"pubsubhub/models"
)

func GetUserById(appConfig *config.ApplicationConfig, userId int) *models.User {
	for _, user := range appConfig.Users {
		if user.Id == userId {
			return user
		}
	}
	return nil
}

func AuthenticateUser(appConfig *config.ApplicationConfig, username string, password string) *models.User {
	for _, user := range appConfig.Users {
		if user.Username == username && user.Password == password {
			return user
		}
	}
	return nil
}
