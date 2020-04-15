package main

import (
	"context"
	"ds-project/common/proto/users"
	"ds-project/common/proto/posts"
	"google.golang.org/grpc"
	"testing"
	"sync"
	"log"
	"fmt"
)

func TestUserNameExists(t *testing.T) {
	fmt.Println("Testing user name exists")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	resp, err := userClient.CheckUserNameExists(context.Background(), &users.GetUserRequest{Username: "varun"})

	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

	log.Println("HERE")
	log.Println(resp.Ok)

}

func TestGetUserExists(t *testing.T) {
	fmt.Println("Testing user name exists")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	resp, err := userClient.GetUsers(context.Background(), &users.GetUsersRequest{})

	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

	log.Println("HERE")
	log.Println(resp.Users)

}


func TestGetUserDetails(t *testing.T) {
	fmt.Println("Testing user name exists")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	resp, err := userClient.GetUser(context.Background(), &users.GetUserRequest{Username: "varun"})

	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

	log.Println("HERE")
	log.Println(resp.Username)

}

func TestMultiplePost(t *testing.T) {
	fmt.Println("Testing add multiple post service")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)
	wg := sync.WaitGroup{}
	for idx := 0; idx < 10000; idx++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()


		response, err := postClient.AddPost(context.Background(), &posts.AddPostRequest{
			Username: "varun",
			Post:     "New POST :" + string(idx) ,
		})

		if err != nil {
			log.Println(err)
			t.Error("fails")
		}
		log.Println(response.Ok)


		}(idx)
	}

	wg.Wait()
}

// func TestTokenValid(t *testing.T){
// 	log.Println("Testing user name exists")
// 	appConfig := config.NewAppConfig()
// 	token := services.GenerateAccessToken(appConfig, "varun")
// 	if !services.CheckAccessTokenValid(appConfig,"varun",token) {
// 		t.Error("Fails")
// 	}
// 	log.Println("token valid")
// }

// func TestGetPostsForUsers(t *testing.T){
// 	log.Println("Testing get post service")
// 	appConfig := config.NewAppConfig()
// 	posts := services.GetPostsForUser(appConfig, "varun")

// 	log.Println(posts[0].Post)

// 	if posts[0].Post != "My name is Varun."{
// 		t.Error("fails")
// 	}
// }

// func TestPostAdd(t *testing.T){
// 	log.Println("Testing add post service")
// 	appConfig := config.NewAppConfig()
// 	services.AddPost(appConfig, "varun", "New POST")

// 	posts := services.GetPostsForUser(appConfig, "varun")
// 	size := len(posts)
// 	log.Println(posts[size-1].Post)

// 	if posts[size-1].Post != "New POST"{
// 		t.Error("fails")
// 	}
// }

// func TestMultiplePost(t *testing.T) {
// 	log.Println("Testing add multiple post service")

// 	appConfig := config.NewAppConfig()
// 	initialPosts := services.GetPostsForUser(appConfig, "varun")
// 	initialPostsLength := len(initialPosts)

// 	wg := sync.WaitGroup{}
// 	for idx := 0; idx < 1000; idx++ {
// 		wg.Add(1)

// 		go func(idx int) {
// 			defer wg.Done()
// 			p := "New POST " + string(idx)
// 			services.AddPost(appConfig, "varun", p)
// 		}(idx)
// 	}

// 	wg.Wait()
// 	finalPosts := services.GetPostsForUser(appConfig, "varun")
// 	finalPostsLength := len(finalPosts)

// 	log.Println(initialPostsLength)
// 	log.Println(finalPostsLength)

// 	if finalPostsLength != initialPostsLength + 1000 {
// 		t.Error("fails")
// 	}
// }

// func TestGetFeedForUsers(t *testing.T){
// 	log.Println("Testing get post service")
// 	appConfig := config.NewAppConfig()
// 	feeds := services.GetFeedForUsername(appConfig, "varun")

// 	log.Println(feeds[0].Post)

// 	if len(feeds) == 0 {
// 		t.Error("fails")
// 	}
// }

// func TestSubscriptions(t *testing.T){
// 	log.Println("Testing get subscription service")
// 	appConfig := config.NewAppConfig()
// 	subscriptions := services.GetSubscriptionsForUsername(appConfig, "varun")

// 	log.Println(subscriptions)

// 	if len(subscriptions) == 0 {
// 		t.Error("fails")
// 	}
// }

// func TestAddSubscriptions(t *testing.T){
// 	log.Println("Testing subscribe service")
// 	appConfig := config.NewAppConfig()
// 	services.Subscribe(appConfig, "varun","ghanu")

// 	subscriptions := services.GetSubscriptionsForUsername(appConfig, "varun")

// 	log.Println(subscriptions)

// 	if subscriptions[len(subscriptions)-1] != "ghanu" {
// 		t.Error("fails")
// 	}
// }

// func TestRemoveSubscriptions(t *testing.T){
// 	log.Println("Testing subscribe service")
// 	appConfig := config.NewAppConfig()
// 	services.Unsubscribe(appConfig, "varun","ghanu")

// 	subscriptions := services.GetSubscriptionsForUsername(appConfig, "varun")
