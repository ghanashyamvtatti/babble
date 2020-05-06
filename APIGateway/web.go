package main

import (
	"ds-project/APIGateway/common"
	"ds-project/APIGateway/controllers"
	"ds-project/common/proto/auth"
	"ds-project/common/proto/posts"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/proto/users"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(controllers.CORSMiddleware())

	// Create connections to various services
	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	authConnection, err := grpc.Dial("localhost:3004", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	subscriptionConnection, err := grpc.Dial("localhost:3005", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// Get clients for connections
	clients := common.ServiceClients{
		UserClient:         users.NewUserServiceClient(userConnection),
		PostClient:         posts.NewPostsServiceClient(postConnection),
		AuthClient:         auth.NewAuthServiceClient(authConnection),
		SubscriptionClient: subscriptions.NewSubscriptionServiceClient(subscriptionConnection),
	}

	authorization := router.Group("/auth")
	{
		authorization.POST("/sign-up", controllers.SignUp(&clients))
		authorization.POST("/sign-in", controllers.SignIn(&clients))
		authorization.POST("/user/:username/sign-out", controllers.SignOut(&clients))
	}

	// Social group
	social := router.Group("/social/user/:username")
	{
		social.Use(controllers.Authenticate(&clients))
		social.GET("/", controllers.GetUserDetails(&clients))
		social.GET("/feed", controllers.GetUserFeed(&clients))
		social.GET("/post", controllers.GetUserPosts(&clients))
		social.POST("/post", controllers.CreatePost(&clients))
		social.POST("/subscribe/:publisher", controllers.Subscribe(&clients))
		social.DELETE("/subscribe/:publisher", controllers.Unsubscribe(&clients))
		social.GET("/subscriptions", controllers.GetSubscriptions(&clients))
	}
	// Get all users
	router.GET("/social/user", controllers.GetUsers(&clients))

	return router

}
