package controllers

import (
	"ds-project/common/proto/models"
	"ds-project/common/proto/users"
	"ds-project/config"
	"ds-project/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
//import (
//	userproto "ds-project-bak/Common/Proto/Users"
//	"ds-project/common/proto/dsl"
//	"ds-project/config"
//	"ds-project/common/proto"
//	"ds-project/dtos"
//	"ds-project/services"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//

func mapUsersToDTO(users map[string]*models.User) []*dtos.User {
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

func GetUsers(clients *config.ServiceClients) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := clients.UserClient.GetUsers(ctx, &users.GetUsersRequest{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, dtos.Response{
				Status:  true,
				Message: "Successfully fetched the data",
				Data: gin.H{
					"result": mapUsersToDTO(response.Users),
				},
			})
		}
	}
}

func GetUserDetails(clients *config.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		username = context.DefaultQuery("username", username)
		response, err := clients.UserClient.GetUser(context, &users.GetUserRequest{Username: username})
		if err != nil {
			context.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, dtos.Response{
				Status:  true,
				Message: "Successfully fetched the data",
				Data: gin.H{
					"user": dtos.User{
						Username: response.Username,
						FullName: response.User.FullName,
					},
				},
			})
		}
	}
}
