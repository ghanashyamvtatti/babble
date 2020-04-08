package services

import (
	"ds-project/DAL"
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
	if _, ok := DAL.GetUser(appConfig, username); ok {
		return true
	}
	return false
}

func GenerateAccessToken(appConfig *config.ApplicationConfig, username string) string {
	token, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	DAL.SetAccessToken(appConfig, username, token.String())
	return token.String()
}

func CheckAccessTokenValid(appConfig *config.ApplicationConfig, username string, token string) bool {
	if t, ok := DAL.GetAccessToken(appConfig, username); ok {
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
	return mapUsersToDTO(DAL.GetUsers(appConfig))
}

func CreateUser(appConfig *config.ApplicationConfig, username string, value models.User) dtos.User {
	DAL.CreateUser(appConfig, username, value)
	appConfig.Users[username] = value

	return dtos.User{
		Username: username,
		FullName: value.FullName,
	}
}

func GetUserByUsername(appConfig *config.ApplicationConfig, username string) dtos.User {
	if user, ok := DAL.GetUser(appConfig, username); ok {
		return dtos.User{
			Username: username,
			FullName: user.FullName,
		}
	} else {
		return dtos.User{}
	}
}

func Login(appConfig *config.ApplicationConfig, username string, password string) bool {
	if user, ok := DAL.GetUser(appConfig, username); ok {
		if utilities.CheckPasswordHash(password, user.Password) {
			return true
		}
	}
	return false
}

func Logout(appConfig *config.ApplicationConfig, username string) {
	if _, ok := DAL.GetAccessToken(appConfig, username); ok {
		DAL.DeleteAccessToken(appConfig, username)
	}
}
