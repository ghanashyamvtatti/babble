package concurrency

import (
	"context"
	"ds-project/common/proto/posts"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"testing"
)

var (
	concurrencyFactor = 1000
)

func TestMultiplePost(t *testing.T) {
	log.Println("Testing Post Services")
	fmt.Println("Testing add multiple post service")

	postConnection, err := grpc.Dial("localhost:3003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	postClient := posts.NewPostsServiceClient(postConnection)
	wg := sync.WaitGroup{}
	for idx := 0; idx < concurrencyFactor; idx++ {
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

		}(idx)
	}

	wg.Wait()

	log.Println("Pass TestMultiplePost")
}
