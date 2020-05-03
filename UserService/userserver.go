package main

import (
	"context"
	"ds-project/UserService/userdal"
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/proto/users"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
	"time"
)

var (
	dialTimeout = 2 * time.Second
)

type UserServer struct {
	users.UnimplementedUserServiceServer
	client *clientv3.Client
	dal    userdal.UserDAL
}

/*
rpc CheckUserNameExists(GetUserRequest) returns (UserExistsResponse);
rpc GetUsers(GetUsersRequest) returns (GetUsersResponse);
rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
rpc GetUser(GetUserRequest) returns (GetUserResponse);
*/

func (server *UserServer) CheckUserNameExists(ctx context.Context, req *users.GetUserRequest) (*users.UserExistsResponse, error) {
	res := make(chan *models.User)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}
	go server.dal.GetUser(request, req.Username, res)

	select {
	case us := <-res:
		if us != nil {
			return &users.UserExistsResponse{Ok: true}, nil
		} else {
			return &users.UserExistsResponse{Ok: false}, nil
		}
	case err := <-request.ErrorChan:
		return &users.UserExistsResponse{Ok: false}, err
	case <-ctx.Done():
		return &users.UserExistsResponse{Ok: false}, ctx.Err()
	}
}

func (server *UserServer) GetUsers(ctx context.Context, req *users.GetUsersRequest) (*users.GetUsersResponse, error) {

	res := make(chan map[string]*models.User)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go server.dal.GetUsers(request, res)

	select {
	case us := <-res:
		return &users.GetUsersResponse{Users: us}, nil
	case err := <-request.ErrorChan:
		return &users.GetUsersResponse{Users: nil}, err
	case <-ctx.Done():
		return &users.GetUsersResponse{Users: nil}, ctx.Err()
	}
}

func (server *UserServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {

	res := make(chan bool)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go server.dal.CreateUser(request, req.Username, req.User, res)

	select {
	case <-res:
		return &users.CreateUserResponse{}, nil
	case err := <-request.ErrorChan:
		return &users.CreateUserResponse{}, err
	case <-ctx.Done():
		return &users.CreateUserResponse{}, ctx.Err()
	}
}

func (server *UserServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {

	res := make(chan *models.User)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    server.client,
		ErrorChan: errorChan,
	}

	go server.dal.GetUser(request, req.Username, res)

	select {
	case r := <-res:
		return &users.GetUserResponse{Username: req.Username, User: r,}, nil
	case err := <-request.ErrorChan:
		return &users.GetUserResponse{Username: req.Username, User: nil,}, err
	case <-ctx.Done():
		return &users.GetUserResponse{Username: req.Username, User: nil,}, ctx.Err()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()

	server := grpc.NewServer()
	users.RegisterUserServiceServer(server, &UserServer{
		client: cli,
		dal:    userdal.UserDAL{Mutex: sync.Mutex{}},
	})
	reflection.Register(server)
	log.Println("User service running on :3002")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
