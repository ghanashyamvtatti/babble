package main

import (
	"context"
	"ds-project/common/proto/dsl"
	"ds-project/common/proto/models"
	"ds-project/common/proto/posts"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sort"
	"sync"
)

type PostsServer struct {
	posts.UnimplementedPostsServiceServer
	mutex     sync.Mutex
	dslClient dsl.DataServiceClient
}

/*
PostService
  rpc AddPost(AddPostRequest) returns (AddPostResponse);
  rpc GetPosts(GetPostsRequest) returns (GetPostsResponse);
  rpc GetFeed(GetPostsRequest) returns (GetPostsResponse);
*/

func (server *PostsServer) AddPost(ctx context.Context, post *posts.AddPostRequest) (*posts.AddPostResponse, error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	dslResponse, err := server.dslClient.AddPost(ctx, &dsl.AddPostRequest{
		Username: post.Username,
		Post:     post.Post,
	})
	if err == nil {
		response := &posts.AddPostResponse{Ok: dslResponse.Ok}
		return response, err
	} else {
		return &posts.AddPostResponse{Ok: false}, err
	}
}

func (server *PostsServer) GetPosts(ctx context.Context, req *posts.GetPostsRequest) (*posts.GetPostsResponse, error) {
	dslResponse, err := server.dslClient.GetPosts(ctx, &dsl.GetPostsRequest{Username: req.Username})
	if err == nil {
		response := &posts.GetPostsResponse{Posts: dslResponse.Posts}
		return response, err
	} else {
		return &posts.GetPostsResponse{
			Posts: nil,
		}, nil
	}
}

func (server *PostsServer) GetFeed(ctx context.Context, req *posts.GetPostsRequest) (*posts.GetPostsResponse, error) {
	subscriptions, _ := server.dslClient.GetSubscriptions(ctx, &dsl.GetSubscriptionsRequest{Username: req.Username})
	var responsePosts []*models.Post

	for _, subscription := range subscriptions.Subscriptions {
		response, _ := server.dslClient.GetPosts(ctx, &dsl.GetPostsRequest{Username: subscription})
		userPosts := response.Posts
		responsePosts = append(responsePosts, userPosts...)
	}
	sort.Slice(responsePosts, func(i, j int) bool {
		iTime, _ := ptypes.Timestamp(responsePosts[i].CreatedAt)
		jTime, _ := ptypes.Timestamp(responsePosts[j].CreatedAt)
		return iTime.After(jTime)
	})
	return &posts.GetPostsResponse{Posts: responsePosts}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Set up a connection to the DSL server.
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	dslClient := dsl.NewDataServiceClient(conn)

	server := grpc.NewServer()
	posts.RegisterPostsServiceServer(server, &PostsServer{dslClient: dslClient})
	reflection.Register(server)
	log.Println("Posts service running on :3003")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
