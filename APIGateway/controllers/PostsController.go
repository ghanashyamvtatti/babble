package controllers

import (
	"ds-project/APIGateway/dtos"
	"ds-project/common"
	"ds-project/common/proto/posts"
	"github.com/gin-gonic/gin"
)

func GetUserFeed(clients *common.ServiceClients) gin.HandlerFunc {
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

func GetUserPosts(clients *common.ServiceClients) gin.HandlerFunc {
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

func CreatePost(clients *common.ServiceClients) gin.HandlerFunc {
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
