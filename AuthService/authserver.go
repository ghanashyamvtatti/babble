package main

import (
	"context"
	"ds-project/AuthService/authdal"
	"ds-project/AuthService/authdal/impl"
	"ds-project/common"
	"ds-project/common/proto/auth"
	"ds-project/common/proto/users"
	"ds-project/common/utilities"
	"github.com/coreos/etcd/clientv3"
	"github.com/google/uuid"
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

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	client     *clientv3.Client
	userClient users.UserServiceClient
	authDAL    authdal.AuthDAL
}

/*
rpc GenerateAccessToken(GenerateTokenRequest) returns (GenerateTokenResponse);
rpc CheckAccessTokenValid(TokenValidityRequest) returns (TokenValidityResponse);
rpc Login(LoginRequest) returns (LoginResponse);
rpc Logout(LogoutRequest) returns (LogoutResponse);
*/

func (s *AuthServer) GenerateAccessToken(ctx context.Context, req *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error) {
	token, err := uuid.NewUUID()

	if err != nil {
		panic(err)
	}

	result := make(chan bool)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    s.client,
		ErrorChan: errorChan,
	}

	go s.authDAL.SetAccessToken(request, req.Username, token.String(), result)

	select {
	case <-result:
		return &auth.GenerateTokenResponse{Token: token.String()}, nil
	case err := <-request.ErrorChan:
		return &auth.GenerateTokenResponse{Token: token.String()}, err
	case <-ctx.Done():
		return &auth.GenerateTokenResponse{Token: token.String()}, ctx.Err()
	}
}

func (s *AuthServer) CheckAccessTokenValid(ctx context.Context, req *auth.TokenValidityRequest) (*auth.TokenValidityResponse, error) {

	result := make(chan string)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    s.client,
		ErrorChan: errorChan,
	}

	go s.authDAL.GetAccessToken(request, req.Username, result)

	select {
	case token := <-result:
		if token == req.Token {
			return &auth.TokenValidityResponse{Ok: true}, nil
		} else {
			return &auth.TokenValidityResponse{Ok: false}, nil
		}
	case err := <-request.ErrorChan:
		return &auth.TokenValidityResponse{Ok: false}, err
	case <-ctx.Done():
		return &auth.TokenValidityResponse{Ok: false}, ctx.Err()
	}

}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	response, err := s.userClient.GetUser(ctx, &users.GetUserRequest{Username: req.Username})

	if err != nil || response.User == nil {
		return &auth.LoginResponse{Ok: false}, err
	} else {
		if utilities.CheckPasswordHash(req.Password, response.User.Password) {
			return &auth.LoginResponse{Ok: true}, nil
		} else {
			return &auth.LoginResponse{Ok: false}, nil
		}
	}
}

func (s *AuthServer) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {

	result := make(chan bool)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    s.client,
		ErrorChan: errorChan,
	}
	go s.authDAL.DeleteAccessToken(request, req.Username, result)

	select {
	case <-result:
		return &auth.LogoutResponse{}, nil
	case err := <-errorChan:
		return &auth.LogoutResponse{}, err
	case <-ctx.Done():
		return &auth.LogoutResponse{}, ctx.Err()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3004")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userConnection, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer userConnection.Close()

	server := grpc.NewServer()
	var authDAL authdal.AuthDAL

	// Below code uses etcd
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()
	authDAL = &impl.EtcdAuthDAL{Mutex: sync.Mutex{}}

	// Below code uses in-memory storage
	/*	appConfig := config.NewAppConfig()
		authDAL = &impl.DSLAuthDAL{Mutex: sync.Mutex{}, AppConfig: appConfig}*/
	
	auth.RegisterAuthServiceServer(server, &AuthServer{
		// Remove the client param if using in-memory storage
		client:     cli,
		userClient: users.NewUserServiceClient(userConnection),
		authDAL:    authDAL,
	})
	reflection.Register(server)
	log.Println("Auth service running on :3004")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
