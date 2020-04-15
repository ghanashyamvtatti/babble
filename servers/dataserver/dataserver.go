package main

import (
	"context"
	"ds-project/DAL"
	"ds-project/common/proto/dsl"
	"ds-project/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type DSLServer struct {
	dsl.UnimplementedDataServiceServer
	appConfig *config.ApplicationConfig
}

/*
Services related to posts
=========================
rpc AddPost(AddPostRequest) returns (AddPostResponse);
rpc GetPosts(GetPostsRequest) returns (GetPostsResponse);
*/
func (server *DSLServer) AddPost(ctx context.Context, post *dsl.AddPostRequest) (*dsl.AddPostResponse, error) {
	response := DAL.AddPost(server.appConfig, post.Username, post.Post)
	return &dsl.AddPostResponse{Ok: response}, nil
}

func (server *DSLServer) GetPosts(ctx context.Context, req *dsl.GetPostsRequest) (*dsl.GetPostsResponse, error) {
	response := DAL.GetPosts(server.appConfig, req.Username)
	return &dsl.GetPostsResponse{Posts: response}, nil
}

/*
Services related to subscriptions
=================================
rpc GetSubscriptions(GetSubscriptionsRequest) returns (GetSubscriptionsResponse);
rpc Subscribe(SubscribeRequest) returns (SubscribeResponse);
rpc Unsubscribe(SubscribeRequest) returns (SubscribeResponse);
*/

func (server *DSLServer) GetSubscriptions(ctx context.Context, req *dsl.GetSubscriptionsRequest) (*dsl.GetSubscriptionsResponse, error) {
	response := DAL.GetSubscriptions(server.appConfig, req.Username)
	return &dsl.GetSubscriptionsResponse{Subscriptions: response}, nil
}

func (server *DSLServer) Subscribe(ctx context.Context, req *dsl.SubscribeRequest) (*dsl.SubscribeResponse, error) {
	response := DAL.Subscribe(server.appConfig, req.Subscriber, req.Publisher)
	return &dsl.SubscribeResponse{Ok: response}, nil
}

func (server *DSLServer) Unsubscribe(ctx context.Context, req *dsl.SubscribeRequest) (*dsl.SubscribeResponse, error) {
	response := DAL.Unsubscribe(server.appConfig, req.Subscriber, req.Publisher)
	return &dsl.SubscribeResponse{Ok: response}, nil
}

/*
Services related to auth
========================
rpc SetAccessToken(SetAccessTokenRequest) returns (UpdateAccessTokenResponse);
rpc GetAccessToken(AccessTokenRequest) returns (GetAccessTokenResponse);
rpc DeleteAccessToken(AccessTokenRequest) returns (UpdateAccessTokenResponse);
*/

func (server *DSLServer) SetAccessToken(ctx context.Context, req *dsl.SetAccessTokenRequest) (*dsl.UpdateAccessTokenResponse, error) {
	response := DAL.SetAccessToken(server.appConfig, req.Username, req.Token)
	return &dsl.UpdateAccessTokenResponse{Ok: response}, nil
}

func (server *DSLServer) GetAccessToken(ctx context.Context, req *dsl.AccessTokenRequest) (*dsl.GetAccessTokenResponse, error) {
	token, ok := DAL.GetAccessToken(server.appConfig, req.Username)
	return &dsl.GetAccessTokenResponse{Token: token, Ok: ok}, nil
}

func (server *DSLServer) DeleteAccessToken(ctx context.Context, req *dsl.AccessTokenRequest) (*dsl.UpdateAccessTokenResponse, error) {
	response := DAL.DeleteAccessToken(server.appConfig, req.Username)
	return &dsl.UpdateAccessTokenResponse{Ok: response}, nil
}

/*
Services related to users
=========================
rpc GetUser(GetUserRequest) returns (GetUserResponse);
rpc GetUsers(GetUsersRequest) returns (GetUsersResponse);
rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
*/

func (server *DSLServer) GetUser(ctx context.Context, req *dsl.GetUserRequest) (*dsl.GetUserResponse, error) {
	user, ok := DAL.GetUser(server.appConfig, req.Username)
	return &dsl.GetUserResponse{
		User: user,
		Ok:   ok,
	}, nil
}

func (server *DSLServer) GetUsers(ctx context.Context, req *dsl.GetUsersRequest) (*dsl.GetUsersResponse, error) {
	response := DAL.GetUsers(server.appConfig)
	return &dsl.GetUsersResponse{Users: response}, nil
}

func (server *DSLServer) CreateUser(ctx context.Context, req *dsl.CreateUserRequest) (*dsl.CreateUserResponse, error) {
	ok := DAL.CreateUser(server.appConfig, req.Username, req.User)
	return &dsl.CreateUserResponse{Ok: ok}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	appConfig := config.NewAppConfig()
	dsl.RegisterDataServiceServer(server, &DSLServer{
		appConfig: appConfig,
	})
	reflection.Register(server)
	log.Println("DSL service running on :3001")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Define all the DAL methods we might ever require (What the services need)

//
