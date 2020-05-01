package main

import (
	"context"
	// "ds-project/common/proto/dsl"
	"ds-project/DAL"
	"ds-project/common/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"github.com/coreos/etcd/clientv3"
    "time"
)

var (  
    dialTimeout    = 2 * time.Second
    requestTimeout = 10 * time.Second
)

type UserServer struct {
	users.UnimplementedUserServiceServer
	kv clientv3.KV
}

/*
rpc CheckUserNameExists(GetUserRequest) returns (UserExistsResponse);
rpc GetUsers(GetUsersRequest) returns (GetUsersResponse);
rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
rpc GetUser(GetUserRequest) returns (GetUserResponse);
*/

func (server *UserServer) CheckUserNameExists(ctx context.Context, req *users.GetUserRequest) (*users.UserExistsResponse, error) {
	_, ok := DAL.GetUser(ctx,server.kv, req.Username)
	return &users.UserExistsResponse{Ok: ok}, nil
}

func (server *UserServer) GetUsers(ctx context.Context, req *users.GetUsersRequest) (*users.GetUsersResponse, error) {
	us := DAL.GetUsers(ctx,server.kv)
	// response, err := server.dslClient.GetUsers(ctx, &dsl.GetUsersRequest{})
	// if err != nil {
	// 	return &users.GetUsersResponse{Users: nil}, err
	// } else {
		return &users.GetUsersResponse{Users: us}, nil
	// }
}

func (server *UserServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	DAL.CreateUser(ctx,server.kv, req.Username, req.User)

	// _, err := server.dslClient.CreateUser(ctx, &dsl.CreateUserRequest{
	// 	Username: req.Username,
	// 	User:     req.User,
	// })
	return &users.CreateUserResponse{}, nil
}

func (server *UserServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	response, _ := DAL.GetUser(ctx,server.kv, req.Username)
	// response, err := server.dslClient.GetUser(ctx, &dsl.GetUserRequest{Username: req.Username})
	// if err != nil {
	// 	return &users.GetUserResponse{}, err
	// } else {
		return &users.GetUserResponse{
			Username: req.Username,
			User:     response,
		}, nil
	// }
}

func main() {
	listener, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
    cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    keyVal := clientv3.NewKV(cli)

	// // Set up a connection to the DSL server.
	// conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(), grpc.WithBlock())

	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// dslClient := dsl.NewDataServiceClient(conn)

	server := grpc.NewServer()
	users.RegisterUserServiceServer(server, &UserServer{kv: keyVal})
	reflection.Register(server)
	log.Println("User service running on :3002")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
