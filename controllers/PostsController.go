package controllers
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