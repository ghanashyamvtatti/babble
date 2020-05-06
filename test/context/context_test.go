package context

import (
	"context"
	"ds-project/common/proto/posts"
	"ds-project/common/proto/subscriptions"
	"ds-project/common/proto/users"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

var (
	contextCancelErrMsg = "rpc error: code = Canceled desc = context canceled"
)

// Test cases for User Service
func TestUserServiceGetUsersWithCancelledContext(t *testing.T) {
	log.Println("Testing UserService with cancelled context")
	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	response, err := userClient.GetUsers(ctx, &users.GetUsersRequest{})
	fmt.Println(response)
	if response == nil && err != nil && err.Error() == "rpc error: code = Canceled desc = context canceled" {
		t.Log("The context cancellation was handled")
	} else {
		t.Fatal("The context cancellation was not handled")
	}
}

// Test cases for SubscriptionService
func
TestGetSubscriptionsContextCancelled(t *testing.T) {
	connection, err := grpc.Dial("localhost:3005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := subscriptions.NewSubscriptionServiceClient(connection)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	response, err := client.GetSubscriptions(ctx, &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	if response != nil || err.Error() != contextCancelErrMsg {
		t.Error("Test case failed")
	}
}

func
TestSubscribeContextCancelled(t *testing.T) {
	connection, err := grpc.Dial("localhost:3005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := subscriptions.NewSubscriptionServiceClient(connection)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.Subscribe(ctx, &subscriptions.SubscribeRequest{Subscriber: "varun", Publisher: "ghanu"})

	// First check err message
	if err == nil || err.Error() != contextCancelErrMsg {
		t.Error("Test case failed")
	}

	// Next check if the subscription is present
	response, _ := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	for _, subscription := range response.Subscriptions {
		if subscription == "ghanu" {
			t.Error("Subscription still exists. Test case failed")
		}
	}
}

func
TestUnsubscribeContextCancelled(t *testing.T) {
	connection, err := grpc.Dial("localhost:3005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := subscriptions.NewSubscriptionServiceClient(connection)

	// Ensure that the subscription exists
	client.Subscribe(context.Background(), &subscriptions.SubscribeRequest{Subscriber: "varun", Publisher: "ghanu"})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.Unsubscribe(ctx, &subscriptions.SubscribeRequest{Subscriber: "varun", Publisher: "ghanu"})

	// First check err message
	if err == nil || err.Error() != contextCancelErrMsg {
		t.Error("Test case failed")
	}

	// Next check if the subscription is present
	response, _ := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	for _, subscription := range response.Subscriptions {
		if subscription == "ghanu" {
			t.Log("Subscription still exists. Test case passed")
			return
		}
	}
	t.Error("Subscription doesn't exist. Test case failed")
}

// Test cases for Posts
func
TestAddPostForContextCancelled(t *testing.T) {
	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	postmsg := "TestAddPostContextCancelled"
	client := posts.NewPostsServiceClient(postConnection)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.AddPost(ctx, &posts.AddPostRequest{Username: "ghanu", Post: postmsg})
	// First check err message
	if err == nil || err.Error() != contextCancelErrMsg {
		t.Error("Test case failed")
	}

	// Next check if the post is present
	response, _ := client.GetPosts(context.Background(), &posts.GetPostsRequest{Username: "ghanu"})
	for _, post := range response.Posts {
		if post.Post == postmsg {
			t.Error("Post still exists. Test case failed")
			return
		}
	}
	t.Log("Test case passed")
}
