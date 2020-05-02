package test

import (
	"context"
	"ds-project/common/proto/auth"
	"ds-project/common/proto/posts"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/proto/users"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"testing"
)

// Test cases for User Service

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

func TestGetUsers(t *testing.T) {
	fmt.Println("Testing get users")

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
	fmt.Println("Testing get user details")

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

// Test cases for Post Service

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
				Post:     "New POST :" + string(idx),
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

func TestGetFeed(t *testing.T) {
	fmt.Println("Testing add multiple post service")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)

	response, err := postClient.GetFeed(context.Background(), &posts.GetPostsRequest{
		Username: "varun",
	})

	if err != nil {
		log.Println(err)
		t.Error("fails")
	}
	log.Println(response.Posts)
}

// Test cases for Auth Service

func TestGenerateAccessToken(t *testing.T) {
	fmt.Println("Testing get user details")

	authConnection, err := grpc.Dial("localhost:3004", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	authClient := auth.NewAuthServiceClient(authConnection)
	resp, err := authClient.GenerateAccessToken(context.Background(), &auth.GenerateTokenRequest{Username: "pratik"})

	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

	log.Println("HERE")
	log.Println(resp.Token)
}

func TestTokenValid(t *testing.T) {
	fmt.Println("Testing access token validity")
	authConnection, err := grpc.Dial("localhost:3004", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	authClient := auth.NewAuthServiceClient(authConnection)
	validityResp, err := authClient.CheckAccessTokenValid(context.Background(), &auth.TokenValidityRequest{Username: "ghanu", Token: "MASTER-TOKEN"})
	if err != nil || !validityResp.Ok {
		log.Println(err)
		t.Error("Fails")
	}
	log.Println("HERE")
	log.Println(validityResp.Ok)
}

// Tests for Subscription Service

func TestSubscriptions(t *testing.T) {
	log.Println("Testing get subscription service")
	connection, err := grpc.Dial("localhost:3005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := subscriptions.NewSubscriptionServiceClient(connection)
	response, err := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}
	log.Println(response.Subscriptions)

	if len(response.Subscriptions) == 0 {
		t.Error("fails")
	}
}

func TestAddSubscriptions(t *testing.T) {
	log.Println("Testing subscribe service")
	connection, err := grpc.Dial("localhost:3005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := subscriptions.NewSubscriptionServiceClient(connection)

	// Create new subscription
	resp, err := client.Subscribe(context.Background(), &subscriptions.SubscribeRequest{
		Subscriber: "varun",
		Publisher:  "ghanu",
	})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

	log.Println(resp)

	// Check subscriptions
	response, err := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}
	log.Println(response.Subscriptions)

	if response.Subscriptions[len(response.Subscriptions)-1] != "ghanu" {
		t.Error("fails")
	}
}

func TestRemoveSubscriptions(t *testing.T) {
	log.Println("Testing unsubscribe service")

	connection, err := grpc.Dial("localhost:3005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := subscriptions.NewSubscriptionServiceClient(connection)

	// Unsubscribe
	resp, err := client.Unsubscribe(context.Background(), &subscriptions.SubscribeRequest{
		Subscriber: "varun",
		Publisher:  "ghanu",
	})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

	log.Println(resp)

	// Check subscriptions
	response, err := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}
	log.Println(response.Subscriptions)

	log.Println(response.Subscriptions)

	if response.Subscriptions[len(response.Subscriptions)-1] == "ghanu" {
		t.Error("fails")
	}
}
