package services

import (
	"ds-project/common/utilities"
	"ds-project/config"
	"ds-project/dtos"
	"ds-project/models"
	"fmt"
	"github.com/google/uuid"
)

func mapUsersToDTO(users map[string]models.User) []*dtos.User {
	var data []*dtos.User
	for username, user := range users {
		userDTO := &dtos.User{
			Username: username,
			FullName: user.FullName,
		}
		data = append(data, userDTO)
	}
	return data
}

func CheckUserNameExists(appConfig *config.ApplicationConfig, username string) bool {
	if _, ok := appConfig.Users[username]; ok {
		return true
	}
	return false
}

func GenerateAccessToken(appConfig *config.ApplicationConfig, username string) string {
	token, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	appConfig.Tokens[username] = token.String()
	return token.String()
}

func CheckAccessTokenValid(appConfig *config.ApplicationConfig, username string, token string) bool {
	if t, ok := appConfig.Tokens[username]; ok {
		if token == t {
			return true
		}
	}
	return false
}

func GetAllUsers(appConfig *config.ApplicationConfig) []*dtos.User {
	token, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
	return mapUsersToDTO(appConfig.Users)
}

func CreateUser(appConfig *config.ApplicationConfig, username string, value models.User) dtos.User {
	appConfig.Users[username] = value

	return dtos.User{
		Username: username,
		FullName: value.FullName,
	}
}

func GetUserByUsername(appConfig *config.ApplicationConfig, username string) dtos.User {
	if CheckUserNameExists(appConfig, username) {
		user := appConfig.Users[username]
		return dtos.User{
			Username: username,
			FullName: user.FullName,
		}
	} else {
		return dtos.User{}
	}
}

func Login(appConfig *config.ApplicationConfig, username string, password string) bool {
	if p, ok := appConfig.Users[username]; ok {
		if utilities.CheckPasswordHash(password, p.Password) {
			return true
		}
	}
	return false
}

func Logout(appConfig *config.ApplicationConfig, username string) {
	if _, ok := appConfig.Tokens[username]; ok {
		delete(appConfig.Tokens, username)
	}
}
