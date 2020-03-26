package controllers

import (
	"ds-project/config"
	"ds-project/dtos"
	"ds-project/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Subscribe(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		publisher := context.Param("publisher")
		if ! services.CheckUserNameExists(appConfig, publisher) {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid publisher",
				Data:    nil,
			})
			return
		}
		services.Subscribe(appConfig, username, publisher)

		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully subscribed",
			Data:    nil,
		})
	}
}

func Unsubscribe(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		publisher := context.Param("publisher")
		if !services.CheckUserNameExists(appConfig, publisher) {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid publisher",
				Data:    nil,
			})
			return
		}
		services.Unsubscribe(appConfig, username, publisher)
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully unsubscribed",
			Data:    nil,
		})
	}
}

func GetSubscriptions(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		subscriptions := services.GetSubscriptionsForUsername(appConfig, username)
		context.JSON(http.StatusOK, dtos.Response{
			Status:  true,
			Message: "Successfully fetched user subscriptions",
			Data: gin.H{
				"subscriptions": subscriptions,
			},
		})
	}
}
