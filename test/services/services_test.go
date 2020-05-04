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
	"testing"
)

// Test cases for User Service

func TestUserNameExists(t *testing.T) {
	log.Println("Testing User Services")

	fmt.Println("Testing user name exists")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	_, er := userClient.CheckUserNameExists(context.Background(), &users.GetUserRequest{Username: "varun"})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	log.Println("Pass TestUserNameExists")

}

func TestGetUsers(t *testing.T) {
	fmt.Println("Testing get users")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)

	_, er := userClient.GetUsers(context.Background(), &users.GetUsersRequest{})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}
	log.Println("Pass TestGetUsers")
}

func TestGetUserDetails(t *testing.T) {
	fmt.Println("Testing get user details")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	_, er := userClient.GetUser(context.Background(), &users.GetUserRequest{Username: "varun"})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	log.Println("Pass TestGetUserDetails")
}

// Test cases for Post Service

func TestGetFeed(t *testing.T) {
	fmt.Println("Testing get feed")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)

	_, er := postClient.GetFeed(context.Background(), &posts.GetPostsRequest{
		Username: "varun",
	})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}
	log.Println("Pass TestGetFeed")
}

// Test cases for Auth Service

func TestGenerateAccessToken(t *testing.T) {
	log.Println("Testing Auth Services")
	fmt.Println("Testing generate token")

	authConnection, err := grpc.Dial("localhost:3004", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	authClient := auth.NewAuthServiceClient(authConnection)
	_, er := authClient.GenerateAccessToken(context.Background(), &auth.GenerateTokenRequest{Username: "pratik"})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	log.Println("Pass TestGenerateAccessToken")
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
	log.Println("Pass TestTokenValid")
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
	_, er := client.Subscribe(context.Background(), &subscriptions.SubscribeRequest{
		Subscriber: "varun",
		Publisher:  "ghanu",
	})
	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	// Check subscriptions
	response, err := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}

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
	if response.Subscriptions[len(response.Subscriptions)-1] == "ghanu" {
		t.Error("fails")
	}
}
