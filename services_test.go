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


func TestGetFeed(t *testing.T) {
	fmt.Println("Testing add multiple post service")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)
	// wg := sync.WaitGroup{}
	// for idx := 0; idx < 10000; idx++ {
	// 	wg.Add(1)

	// 	go func(idx int) {
	// 		defer wg.Done()


		response, err := postClient.GetFeed(context.Background(), &posts.GetPostsRequest{
			Username: "varun",
		})

		if err != nil {
			log.Println(err)
			t.Error("fails")
		}
		log.Println(response.Posts)


	// 	}(idx)
	// }

	// wg.Wait()
}
