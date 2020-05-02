package controllers

import (
	"ds-project/APIGateway/dtos"
	"ds-project/common"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/proto/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Subscribe(clients *common.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		publisher := context.Param("publisher")
		response, err := clients.UserClient.CheckUserNameExists(context, &users.GetUserRequest{Username: username})
		if err != nil || ! response.Ok {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid publisher",
				Data:    nil,
			})
			return
		}
		_, subscriptionError := clients.SubscriptionClient.Subscribe(context, &subscriptions.SubscribeRequest{
			Subscriber: username,
			Publisher:  publisher,
		})
		if subscriptionError != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: subscriptionError.Error(),
				Data:    nil,
			})
			return
		}

		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully subscribed",
			Data:    nil,
		})
	}
}

func Unsubscribe(clients *common.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		publisher := context.Param("publisher")
		response, err := clients.UserClient.CheckUserNameExists(context, &users.GetUserRequest{Username: username})
		if err != nil || ! response.Ok {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid publisher",
				Data:    nil,
			})
			return
		}
		_, subscriptionError := clients.SubscriptionClient.Unsubscribe(context, &subscriptions.SubscribeRequest{
			Subscriber: username,
			Publisher:  publisher,
		})
		if subscriptionError != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: subscriptionError.Error(),
				Data:    nil,
			})
			return
		}
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully unsubscribed",
			Data:    nil,
		})
	}
}

func GetSubscriptions(clients *common.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		response, err := clients.SubscriptionClient.GetSubscriptions(context, &subscriptions.GetSubscriptionsRequest{Username: username})
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		context.JSON(http.StatusOK, dtos.Response{
			Status:  true,
			Message: "Successfully fetched user subscriptions",
			Data: gin.H{
				"subscriptions": response.Subscriptions,
			},
		})
	}
}
