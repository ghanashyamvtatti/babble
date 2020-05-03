package context

import (
	"context"
	"ds-project/UserService/userdal"
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/proto/subscriptions"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"log"
	"sync"
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

func TestUserDALStorageGetUsers(t *testing.T) {
	log.Println("Testing User DAL Storage")
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
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

	select {
	case us := <-res:
		fmt.Println(us)
	case err := <-errorChan:
		fmt.Println(err)
		t.Error("fails")
	case <-ctx.Done():
		fmt.Println(ctx.Done())
		t.Error("fails")
	}
}

func TestUserDALStorageGetUsersWithCancelledContext(t *testing.T) {
	log.Println("Testing User DAL Storage with Cancelled Context")
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
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

	select {
	case us := <-res:
		fmt.Println(us)
	case err := <-errorChan:
		fmt.Println(err)
		t.Error("fails")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("Cancelled context case")
		// fmt.Println(ctx.Done())
		// t.Error("fails")
	}

	//defer cancel()
}

// Test cases for SubscriptionService

func TestGetSubscriptionsContextCancelled(t *testing.T) {
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

func TestSubscribeContextCancelled(t *testing.T) {
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

func TestUnsubscribeContextCancelled(t *testing.T) {
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
