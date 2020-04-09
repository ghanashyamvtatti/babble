package DAL

import (
	"ds-project/common/proto/models"
	"ds-project/config"
)

func GetUser(appConfig *config.ApplicationConfig, username string) (*models.User, bool) {
	user, ok := appConfig.Users[username]
	return user, ok
}

func GetUsers(appConfig *config.ApplicationConfig) map[string]*models.User {
	return appConfig.Users
}

func CreateUser(appConfig *config.ApplicationConfig, username string, value *models.User) bool {
	appConfig.Users[username] = value
	return true
}
