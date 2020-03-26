package controllers

import (
	"ds-project/config"
	"ds-project/dtos"
	"ds-project/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dtos.Response{
			Status:  true,
			Message: "Successfully fetched the data",
			Data: gin.H{
				"result": services.GetAllUsers(appConfig),
			},
		})
	}
}

func GetUserDetails(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		user := services.GetUserByUsername(appConfig, username)
		context.JSON(http.StatusOK, dtos.Response{
			Status:  true,
			Message: "Successfully fetched the data",
			Data: gin.H{
				"user": user,
			},
		})
	}
}
