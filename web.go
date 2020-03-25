package main

import (
	"github.com/gin-gonic/gin"
	"pubsubhub/config"
	"pubsubhub/controllers"
)

func main() {

	router := gin.Default()

	appConfig := config.NewAppConfig()

	// Auth group
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp(appConfig))
		auth.POST("/sign-in", controllers.SignIn(appConfig))
		auth.POST("/sign-out", controllers.SignOut())
	}

	// Social group
	social := router.Group("/social/user/:userId")
	{
		social.GET("/", controllers.GetUserDetails(appConfig))
		social.GET("/feed", controllers.GetUserFeed(appConfig))
		social.GET("/post", controllers.GetUserPosts(appConfig))
		social.POST("/post", controllers.CreatePost(appConfig))
		social.POST("/subscribe/:pubUserId", controllers.Subscribe(appConfig))
		social.DELETE("/subscribe/:pubUserId", controllers.Unsubscribe(appConfig))
	}

	router.Run(":8080")
}
