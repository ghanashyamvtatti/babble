package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"pubsubhub/config"
	"pubsubhub/dtos"
	"strconv"
)

var appConfig config.ApplicationConfig

func main() {

	router := gin.Default()

	appConfig := config.NewAppConfig()

	// Auth group
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", func(context *gin.Context) {
			var user dtos.User
			if context.ShouldBind(&user) == nil {
				appConfig.AddUser(user.FullName, user.Username, user.Password)
				context.String(200, "User successfully created")
			} else {
				context.String(500, "Unable to create user")
			}
		})
		auth.POST("/sign-in", func(context *gin.Context) {
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
		})
		auth.POST("/sign-out", func(context *gin.Context) {
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully signed out the user",
				Data:    render.JSON{},
			})
		})
	}

	// Social group
	social := router.Group("/social/user/:userId")
	{
		social.GET("/", func(context *gin.Context) {
			userId, err := strconv.Atoi(context.Param("userId"))
			if err != nil {
				context.JSON(500, dtos.Response{
					Status:  false,
					Message: "Invalid userId",
					Data:    render.JSON{},
				})
				return
			}
			user := appConfig.GetUserById(userId)
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
		})
		social.GET("/feed", func(context *gin.Context) {
			userId, err := strconv.Atoi(context.Param("userId"))
			if err != nil {
				context.JSON(500, dtos.Response{
					Status:  false,
					Message: "Invalid userId",
					Data:    render.JSON{},
				})
				return
			}
			feed := appConfig.GetFeedForUserId(userId)
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully fetched feed",
				Data: render.JSON{
					Data: feed,
				},
			})
		})
		social.GET("/post", func(context *gin.Context) {
			userId, err := strconv.Atoi(context.Param("userId"))
			if err != nil {
				context.JSON(500, dtos.Response{
					Status:  false,
					Message: "Invalid userId",
					Data:    render.JSON{},
				})
				return
			}
			posts := appConfig.GetPostsForUserId(userId)
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully fetched user posts",
				Data: render.JSON{
					Data: posts,
				},
			})
		})
		social.POST("/post", func(context *gin.Context) {
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
				appConfig.AddPost(userId, post.Post)
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
		})
		social.POST("/subscribe/:pubUserId", func(context *gin.Context) {
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
			appConfig.Subscribe(userId, publisherUserId)

			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully subscribed",
				Data:    render.JSON{},
			})
		})
		social.DELETE("/subscribe/:pubUserId", func(context *gin.Context) {
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
			appConfig.Unsubscribe(userId, publisherUserId)
			context.JSON(200, dtos.Response{
				Status:  true,
				Message: "Successfully unsubscribed",
				Data:    render.JSON{},
			})
		})
	}

	router.Run(":8080")
}
