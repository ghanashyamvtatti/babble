package context

import (
	"context"
	"ds-project/common/proto/users"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
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
