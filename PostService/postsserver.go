package main

import (
	"context"
	"ds-project/PostService/postdal"
	"ds-project/PostService/postdal/impl"
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
	"ds-project/common/proto/subscriptions"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
	"time"
)

type PostsServer struct {
	posts.UnimplementedPostsServiceServer
	client             *clientv3.Client
	subscriptionClient subscriptions.SubscriptionServiceClient
	postDAL            postdal.PostDAL
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

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go server.postDAL.AddPost(request, post.Username, post.Post, result)

	select {
	case res := <-result:
		return &posts.AddPostResponse{Ok: res}, nil
	case err := <-errorChan:
		return &posts.AddPostResponse{Ok: false}, err
	case <-ctx.Done():
		res := <-result
		if res {
			delRes := make(chan bool)
			req := common.DALRequest{
				Ctx:       context.Background(),
				Client:    server.client,
				ErrorChan: errorChan,
			}
			go server.postDAL.DeletePost(req, post.Username, post.Post, delRes)
		}
		return &posts.AddPostResponse{Ok: false}, ctx.Err()
	}
}

func (server *PostsServer) DeletePost(ctx context.Context, post *posts.AddPostRequest) (*posts.AddPostResponse, error) {
	result := make(chan bool)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go server.postDAL.DeletePost(request, post.Username, post.Post, result)

	select {
	case res := <-result:
		return &posts.AddPostResponse{Ok: res}, nil
	case err := <-errorChan:
		return &posts.AddPostResponse{Ok: false}, err
	case <-ctx.Done():
		res := <-result
		if res {
			addRes := make(chan bool)
			req := common.DALRequest{
				Ctx:       context.Background(),
				Client:    server.client,
				ErrorChan: errorChan,
			}
			go server.postDAL.AddPost(req, post.Username, post.Post, addRes)
		}
		return &posts.AddPostResponse{Ok: false}, ctx.Err()
	}
}

func (server *PostsServer) GetPosts(ctx context.Context, req *posts.GetPostsRequest) (*posts.GetPostsResponse, error) {
	result := make(chan *posts.GetPostsResponse)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go server.postDAL.GetPosts(request, req.Username, result)

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

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}
	subs, _ := server.subscriptionClient.GetSubscriptions(ctx, &subscriptions.GetSubscriptionsRequest{Username: req.Username})
	go server.postDAL.GetFeed(request, subs.Subscriptions, result)

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

	subscriptionConnection, err := grpc.Dial("localhost:3005", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer subscriptionConnection.Close()

	server := grpc.NewServer()
	var postDAL postdal.PostDAL

	// Below code uses etcd
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()
	postDAL = &impl.EtcdPostDAL{Mutex: sync.Mutex{}}

	// Below code uses in-memory storage
	/*	appConfig := config.NewAppConfig()
		postDAL = &impl.DSLPostDAL{
			Mutex:     sync.Mutex{},
			AppConfig: appConfig,
		}*/

	posts.RegisterPostsServiceServer(server, &PostsServer{
		// Remove the client param if using in-memory storage
		client:             cli,
		subscriptionClient: subscriptions.NewSubscriptionServiceClient(subscriptionConnection),
		postDAL:            postDAL,
	})
	reflection.Register(server)
	log.Println("Posts service running on :3003")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
