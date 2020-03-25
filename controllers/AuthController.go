package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"pubsubhub/config"
	"pubsubhub/dtos"
)

func SignUp(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user dtos.User
		if context.ShouldBind(&user) == nil {
			appConfig.AddUser(user.FullName, user.Username, user.Password)
			context.String(200, "User successfully created")
		} else {
			context.String(500, "Unable to create user")
		}
	}
}

func SignIn(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user dtos.User
		if context.ShouldBind(&user) == nil {
			authenticatedUser := appConfig.AuthenticateUser(user.Username, user.Password)
			if authenticatedUser != nil {
				context.JSON(200, dtos.Response{Status: true, Message: "Sign in successful", Data: render.JSON{Data: authenticatedUser}})
				return
			}
		}
		context.JSON(500, dtos.Response{
			Status:  false,
			Message: "Invalid username or password",
			Data:    render.JSON{},
		})
	}
}

func SignOut() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully signed out the user",
			Data:    render.JSON{},
		})
	}
}
