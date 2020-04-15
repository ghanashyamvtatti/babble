package controllers

import (
	"ds-project/common/proto/posts"
	//"ds-project/common/proto/models"
	//"ds-project/common/proto/users"
	//"ds-project/common/utilities"
	"ds-project/config"
	"ds-project/dtos"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func GetUserFeed(clients *config.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {

		response, err := clients.PostClient.GetFeed(context, &posts.GetPostsRequest{
			Username: context.Param("username"),
		})

		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid username",
				Data:    nil,
			})
			context.Abort()
		}

		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully fetched user posts",
			Data: gin.H{
				"feed": response.Posts,
			},
		})
	}
}

func GetUserPosts(clients *config.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		username = context.DefaultQuery("username", username)
		response, err := clients.PostClient.GetPosts(context, &posts.GetPostsRequest{
			Username: username,
		})

		if err != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid username",
				Data:    nil,
			})
			context.Abort()
		}

		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully fetched user posts",
			Data: gin.H{
				"posts": response.Posts,
			},
		})
	}
}

func CreatePost(clients *config.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {

		var post dtos.Post
		if context.ShouldBind(&post) != nil {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid request json",
				Data:    nil,
			})
			context.Abort()
		}

		response, err := clients.PostClient.AddPost(context, &posts.AddPostRequest{
			Username: context.Param("username"),
			Post:     post.Post,
		})

		if err == nil && response.Ok {

			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully added post",
				Data:    nil,
			})
		} else {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "add post failure",
				Data:    nil,
			})
		}
	}
}

/*
import (
	"ds-project/config"
	"ds-project/dtos"
	"github.com/gin-gonic/gin"
)

func GetUserFeed(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		feed := services.GetFeedForUsername(appConfig, username)
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully fetched feed",
			Data: gin.H{
				"feed": feed,
			},
		})
	}
}

func GetUserPosts(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		posts := services.GetPostsForUser(appConfig, username)
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully fetched user posts",
			Data: gin.H{
				"posts": posts,
			},
		})
	}
}

func CreatePost(appConfig *config.ApplicationConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		var post dtos.Post
		if context.ShouldBind(&post) == nil {
			services.AddPost(appConfig, username, post.Post)
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully added post",
				Data:    nil,
			})
			return
		}
		context.JSON(500, dtos.Response{
			Status:  false,
			Message: "Unable to add post",
			Data:    nil,
		})
	}
}
*/
