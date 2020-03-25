package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"pubsubhub/config"
	"pubsubhub/dtos"
	"pubsubhub/services"
	"strconv"
)

func GetUserDetails(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("userId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid userId",
				Data:    render.JSON{},
			})
			return
		}
		user := services.GetUserById(appConfig, userId)
		if user != nil {
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "",
				Data: render.JSON{
					Data: user,
				},
			})
		} else {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Unable to find user given userId",
				Data:    render.JSON{},
			})
		}
	}
}

func GetUserFeed(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("userId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid userId",
				Data:    render.JSON{},
			})
			return
		}
		feed := services.GetFeedForUserId(appConfig, userId)
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully fetched feed",
			Data: render.JSON{
				Data: feed,
			},
		})
	}
}

func GetUserPosts(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("userId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid userId",
				Data:    render.JSON{},
			})
			return
		}
		posts := services.GetPostsForUserId(appConfig, userId)
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully fetched user posts",
			Data: render.JSON{
				Data: posts,
			},
		})
	}
}

func CreatePost(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("userId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid userId",
				Data:    render.JSON{},
			})
			return
		}
		var post dtos.Post
		if context.ShouldBind(&post) == nil {
			services.AddPost(appConfig, userId, post.Post)
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully added post",
				Data:    render.JSON{},
			})
			return
		}
		context.JSON(500, dtos.Response{
			Status:  false,
			Message: "Unable to add post",
			Data:    render.JSON{},
		})
	}
}

func Subscribe(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("userId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid userId",
				Data:    render.JSON{},
			})
			return
		}
		publisherUserId, err := strconv.Atoi(context.Param("pubUserId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid publisher ID",
				Data:    render.JSON{},
			})
			return
		}
		services.Subscribe(appConfig, userId, publisherUserId)

		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully subscribed",
			Data:    render.JSON{},
		})
	}
}

func Unsubscribe(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := strconv.Atoi(context.Param("userId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid userId",
				Data:    render.JSON{},
			})
			return
		}
		publisherUserId, err := strconv.Atoi(context.Param("pubUserId"))
		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid publisher ID",
				Data:    render.JSON{},
			})
			return
		}
		services.Unsubscribe(appConfig, userId, publisherUserId)
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully unsubscribed",
			Data:    render.JSON{},
		})
	}
}
