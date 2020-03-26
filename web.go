package main

import (
	"ds-project/config"
	"ds-project/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	appConfig := config.NewAppConfig()

	// Auth group
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp(appConfig))
		auth.POST("/sign-in", controllers.SignIn(appConfig))
		auth.POST("/user/:username/sign-out", controllers.SignOut(appConfig))
	}

	// Social group
	social := router.Group("/social/user/:username")
	{
		social.Use(controllers.Authenticate(appConfig))
		social.GET("/", controllers.GetUserDetails(appConfig))
		social.GET("/feed", controllers.GetUserFeed(appConfig))
		social.GET("/post", controllers.GetUserPosts(appConfig))
		social.POST("/post", controllers.CreatePost(appConfig))
		social.POST("/subscribe/:publisher", controllers.Subscribe(appConfig))
		social.DELETE("/subscribe/:publisher", controllers.Unsubscribe(appConfig))
		social.GET("/subscriptions", controllers.GetSubscriptions(appConfig))
	}

	router.Run(":8080")
}
