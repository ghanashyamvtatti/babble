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
	"time"
	"github.com/coreos/etcd/clientv3"
	"ds-project/UserService/userdal"
	"ds-project/common/proto/models"
	"ds-project/common"
)


var (  
    dialTimeout    = 2 * time.Second
    requestTimeout = 10 * time.Second
)
// Test cases for User Service

func TestUserDALStorageGetUsers(t *testing.T) {
	log.Println("Testing User DAL Storage")
	cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    // keyVal := clientv3.NewKV(cli)

	res := make(chan *models.User)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       context.Background(),
		Client:    cli,
		ErrorChan: errorChan,
	}

	ctx := context.Background()
	dal := userdal.UserDAL{Mutex: sync.Mutex{}}

	go dal.GetUser(request, "ghanu", res)

	select{
	case us := <-res:
		fmt.Println(us)
	case err := <-errorChan:
		fmt.Println(err)
		t.Error("fails")
	case <- ctx.Done():
		fmt.Println(ctx.Done())
		t.Error("fails")
	}
}

func TestUserDALStorageGetUsersWithCancelledContext(t *testing.T) {
	log.Println("Testing User DAL Storage with Cancelled Context")
	cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    // keyVal := clientv3.NewKV(cli)

	res := make(chan *models.User)
	errorChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	//ctx, _ := context.WithTimeout(context.Background(),time.Duration(2*time.Second))

    


	request := common.DALRequest{
		Ctx:       ctx,
		Client:    cli,
		ErrorChan: errorChan,
	}

	
	dal := userdal.UserDAL{Mutex: sync.Mutex{}}

	// time.Sleep(3 * time.Second)
	go dal.GetUser(request, "ghanu", res)


	select{
	case us := <-res:
		fmt.Println(us)
	case err := <-errorChan:
		fmt.Println(err)
		t.Error("fails")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("Cancelled context case")
		// fmt.Println(ctx.Done())
		// t.Error("fails")
	}

	//defer cancel()
}


// Test cases for User Service

func TestUserNameExists(t *testing.T) {
	log.Println("Testing User Services")

	fmt.Println("Testing user name exists")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	// resp, err := userClient.CheckUserNameExists(context.Background(), &users.GetUserRequest{Username: "varun"})
	_, er := userClient.CheckUserNameExists(context.Background(), &users.GetUserRequest{Username: "varun"})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	log.Println("Pass TestUserNameExists")
	// log.Println(resp.Ok)

}

func TestGetUsers(t *testing.T) {
	fmt.Println("Testing get users")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	// resp, err := userClient.GetUsers(context.Background(), &users.GetUsersRequest{})

	_, er := userClient.GetUsers(context.Background(), &users.GetUsersRequest{})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}
	log.Println("Pass TestGetUsers")

	// log.Println("HERE")
	// log.Println(resp.Users)

}

func TestGetUserDetails(t *testing.T) {
	fmt.Println("Testing get user details")

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := users.NewUserServiceClient(userConnection)
	// resp, err := userClient.GetUser(context.Background(), &users.GetUserRequest{Username: "varun"})
	_, er := userClient.GetUser(context.Background(), &users.GetUserRequest{Username: "varun"})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	log.Println("Pass TestGetUserDetails")



	// log.Println("HERE")
	// log.Println(resp.Username)
}

// Test cases for Post Service

func TestMultiplePost(t *testing.T) {
	log.Println("Testing Post Services")
	fmt.Println("Testing add multiple post service")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)
	wg := sync.WaitGroup{}
	for idx := 0; idx < 1000; idx++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			_, err := postClient.AddPost(context.Background(), &posts.AddPostRequest{
				Username: "varun",
				Post:     "New POST :" + string(idx),
			})

			if err != nil {
				log.Println(err)
				t.Error("fails")
			}
			// log.Println(response.Ok)

		}(idx)
	}

	wg.Wait()

	log.Println("Pass TestMultiplePost")
}

func TestGetFeed(t *testing.T) {
	fmt.Println("Testing get feed")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)

	// response, err := postClient.GetFeed(context.Background(), &posts.GetPostsRequest{
	// 	Username: "varun",
	// })

	_, er := postClient.GetFeed(context.Background(), &posts.GetPostsRequest{
		Username: "varun",
	})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}
	// log.Println(response.Posts)
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
	// resp, err := authClient.GenerateAccessToken(context.Background(), &auth.GenerateTokenRequest{Username: "pratik"})

	_, er := authClient.GenerateAccessToken(context.Background(), &auth.GenerateTokenRequest{Username: "pratik"})

	if er != nil {
		log.Println(er)
		t.Error("fails")
	}

	log.Println("Pass TestGenerateAccessToken")
	// log.Println(resp.Token)
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
	// log.Println("HERE")
	// log.Println(validityResp.Ok)
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
	// log.Println(response.Subscriptions)

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

	// log.Println(resp)

	// Check subscriptions
	response, err := client.GetSubscriptions(context.Background(), &subscriptions.GetSubscriptionsRequest{Username: "varun"})
	if err != nil {
		log.Println(err)
		t.Error("fails")
	}
	// log.Println(response.Subscriptions)

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
	// log.Println(response.Subscriptions)

	// log.Println(response.Subscriptions)

	if response.Subscriptions[len(response.Subscriptions)-1] == "ghanu" {
		t.Error("fails")
	}
}
