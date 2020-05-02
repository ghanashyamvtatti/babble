package main

import (
	"context"
	"ds-project/DAL/postdal"
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
	subscriptions "ds-project/common/proto/subscriptions"
	"ds-project/config"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type PostsServer struct {
	posts.UnimplementedPostsServiceServer
	client             *clientv3.Client
	subscriptionClient subscriptions.SubscriptionServiceClient
}

var (
	dialTimeout = 2 * time.Second
)

/*
PostService
  rpc AddPost(AddPostRequest) returns (AddPostResponse);
  rpc GetPosts(GetPostsRequest) returns (GetPostsResponse);
  rpc GetFeed(GetPostsRequest) returns (GetPostsResponse);
*/

func (server *PostsServer) AddPost(ctx context.Context, post *posts.AddPostRequest) (*posts.AddPostResponse, error) {
	result := make(chan bool)
	errorChan := make(chan error)

	request := config.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go postdal.AddPost(request, post.Username, post.Post, result)

	select {
	case res := <-result:
		return &posts.AddPostResponse{Ok: res}, nil
	case err := <-errorChan:
		return &posts.AddPostResponse{Ok: false}, err
	case <-ctx.Done():
		return &posts.AddPostResponse{Ok: false}, ctx.Err()
	}
}

func (server *PostsServer) GetPosts(ctx context.Context, req *posts.GetPostsRequest) (*posts.GetPostsResponse, error) {
	result := make(chan *posts.GetPostsResponse)
	errorChan := make(chan error)

	request := config.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go postdal.GetPosts(request, req.Username, result)

	select {
	case res := <-result:
		return res, nil
	case err := <-errorChan:
		return &posts.GetPostsResponse{Posts: nil}, err
	case <-ctx.Done():
		return &posts.GetPostsResponse{}, ctx.Err()
	}
}

func (server *PostsServer) GetFeed(ctx context.Context, req *posts.GetPostsRequest) (*posts.GetPostsResponse, error) {
	result := make(chan []*models.Post)
	errorChan := make(chan error)

	request := config.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}
	subs, _ := server.subscriptionClient.GetSubscriptions(ctx, &subscriptions.GetSubscriptionsRequest{Username: req.Username})
	go postdal.GetFeed(request, subs.Subscriptions, result)

	select {
	case responsePosts := <-result:
		return &posts.GetPostsResponse{Posts: responsePosts}, nil
	case err := <-errorChan:
		return &posts.GetPostsResponse{Posts: nil}, err
	case <-ctx.Done():
		return &posts.GetPostsResponse{Posts: nil}, ctx.Err()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Set up a connection to etcd.
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()

	subscriptionConnection, err := grpc.Dial("localhost:3005", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	posts.RegisterPostsServiceServer(server, &PostsServer{client: cli, subscriptionClient: subscriptions.NewSubscriptionServiceClient(subscriptionConnection)})
	reflection.Register(server)
	log.Println("Posts service running on :3003")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
