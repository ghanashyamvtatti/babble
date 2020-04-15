package main

import (
	"context"
	"ds-project/common/proto/dsl"
	"ds-project/common/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type UserServer struct {
	users.UnimplementedUserServiceServer
	dslClient dsl.DataServiceClient
}

/*
rpc CheckUserNameExists(GetUserRequest) returns (UserExistsResponse);
rpc GetUsers(GetUsersRequest) returns (GetUsersResponse);
rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
rpc GetUser(GetUserRequest) returns (GetUserResponse);
*/

func (server *UserServer) CheckUserNameExists(ctx context.Context, req *users.GetUserRequest) (*users.UserExistsResponse, error) {
	response, err := server.dslClient.GetUser(ctx, &dsl.GetUserRequest{Username: req.Username})
	if err != nil {
		return &users.UserExistsResponse{Ok: false}, err
	} else {
		return &users.UserExistsResponse{Ok: response.Ok}, err
	}
}

func (server *UserServer) GetUsers(ctx context.Context, req *users.GetUsersRequest) (*users.GetUsersResponse, error) {
	response, err := server.dslClient.GetUsers(ctx, &dsl.GetUsersRequest{})
	if err != nil {
		return &users.GetUsersResponse{Users: nil}, err
	} else {
		return &users.GetUsersResponse{Users: response.Users}, err
	}
}

func (server *UserServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	_, err := server.dslClient.CreateUser(ctx, &dsl.CreateUserRequest{
		Username: req.Username,
		User:     req.User,
	})
	return &users.CreateUserResponse{}, err
}

func (server *UserServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	response, err := server.dslClient.GetUser(ctx, &dsl.GetUserRequest{Username: req.Username})
	if err != nil {
		return &users.GetUserResponse{}, err
	} else {
		return &users.GetUserResponse{
			Username: req.Username,
			User:     response.User,
		}, err
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3002")
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
	users.RegisterUserServiceServer(server, &UserServer{dslClient: dslClient})
	reflection.Register(server)
	log.Println("User service running on :3002")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
