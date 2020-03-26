package controllers

import (
	"ds-project/common/utilities"
	"ds-project/config"
	"ds-project/dtos"
	"ds-project/models"
	"ds-project/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login dtos.UserLogin

		if err := ctx.ShouldBind(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.Response{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		if services.Login(appConfig, login.Username, login.Password) {
			ctx.JSON(http.StatusOK, dtos.Response{
				Status:  true,
				Message: "Successfully logged in",
				Data: gin.H{
					"user":  login.Username,
					"token": services.GenerateAccessToken(appConfig, login.Username),
				},
			})
		} else {

			ctx.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: "Authentication failed",
				Data:    nil,
			})
		}
	}
}

func SignUp(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var registerUser dtos.UserRegistration
		if err := ctx.ShouldBind(&registerUser); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.Response{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		if !services.CheckUserNameExists(appConfig, registerUser.Username) {
			password, err := utilities.HashPassword(registerUser.Password)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, dtos.Response{
					Status:  false,
					Message: "Unable to create user",
					Data:    nil,
				})
				return
			}
			services.CreateUser(appConfig, registerUser.Username, models.User{FullName: registerUser.FullName, Password: password})
		}

		ctx.JSON(http.StatusOK, dtos.Response{
			Status:  true,
			Message: "Successfully registered user",
			Data: gin.H{
				"user":  registerUser.Username,
				"token": services.GenerateAccessToken(appConfig, registerUser.Username),
			},
		})
	}
}

func SignOut(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		services.Logout(appConfig, context.Param("username"))
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully signed out the user",
			Data:    nil,
		})
	}
}

func Authenticate(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		if ! services.CheckUserNameExists(appConfig, username) {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid username",
				Data:    nil,
			})
			context.Abort()
		}
		if !services.CheckAccessTokenValid(appConfig, username, context.GetHeader("token")) {
			context.JSON(http.StatusUnauthorized, dtos.Response{
				Status:  false,
				Message: "Missing or invalid access token",
				Data:    nil,
			})
			context.Abort()
		}
	}
}
