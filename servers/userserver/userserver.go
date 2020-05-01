package main

import (
	"context"
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
	result := make(chan *models.User)
	errorChan := make(chan error)
	go DAL.GetUser(ctx,server.kv, req.Username,result, errorChan)

	select{
	case <- us := result:
		if us.Username == "" {
			return &users.UserExistsResponse{Ok: true}, nil
		}else{
			return &users.UserExistsResponse{Ok: false}, nil
		}
	case <- err := errorChan:
		return &users.UserExistsResponse{Ok: false}, err
	case <- ctx.Done():
		return &users.UserExistsResponse{Ok: false}, ctx.Err()
	}
}

func (server *UserServer) GetUsers(ctx context.Context, req *users.GetUsersRequest) (*users.GetUsersResponse, error) {

	result := make(chan map[string]*models.User)
	errorChan := make(chan error)

	go DAL.GetUsers(ctx,server.kv,result, errorChan)

	select{
	case <- us := result:
		return &users.GetUsersResponse{Users: us}, nil
	case <- err := errorChan:
		return &users.GetUsersResponse{Users: nil}, err
	case <- ctx.Done():
		return &users.GetUsersResponse{Users: nil}, ctx.Err()
	}
}

func (server *UserServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {

	result := make(chan bool)
	errorChan := make(chan error)

	go DAL.CreateUser(ctx,server.kv, req.Username, req.User,result, errorChan)

	select{
	case <- result:
		return &users.CreateUserResponse{}, nil
	case <- err := errorChan:
		return &users.CreateUserResponse{}, err
	case <- ctx.Done():
		return &users.CreateUserResponse{}, ctx.Err()
	}
}

func (server *UserServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	
	result := make(chan *models.User)
	errorChan := make(chan error)
	go DAL.GetUser(ctx,server.kv, req.Username,result, errorChan)

	select{
	case <- r := result:
		return &users.GetUserResponse{Username: req.Username,User:r,}, nil
	case <- err := errorChan:
		return &users.GetUserResponse{Username: req.Username,User:nil,}, err
	case <- ctx.Done():
		return &users.GetUserResponse{Username: req.Username,User:response,}, ctx.Err()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    keyVal := clientv3.NewKV(cli)

	server := grpc.NewServer()
	users.RegisterUserServiceServer(server, &UserServer{kv: keyVal})
	reflection.Register(server)
	log.Println("User service running on :3002")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
