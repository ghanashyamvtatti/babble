package main

import (
	"context"
	"ds-project/common/proto/auth"
	"ds-project/common/utilities"
	"ds-project/DAL"
	"github.com/google/uuid"
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

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	kv clientv3.KV
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

	result := make(chan string)
	errorChan := make(chan error)

	go DAL.SetAccessToken(ctx,s.kv, req.Username, token.String(), result, errorChan)

	select{
	case <- result:
		return &auth.GenerateTokenResponse{Token: token.String()}, nil
	case <- err := errorChan:
		return &auth.GenerateTokenResponse{Token: token.String()}, err
	case <- ctx.Done():
		return &auth.GenerateTokenResponse{Token: token.String()}, ctx.Err(
	}
}

func (s *AuthServer) CheckAccessTokenValid(ctx context.Context, req *auth.TokenValidityRequest) (*auth.TokenValidityResponse, error) {

	result := make(chan string)
	errorChan := make(chan error)

	go DAL.GetAccessToken(ctx,s.kv, req.Username, result, errorChan)

	select{
	case <-token := result:
		if token == req.Token {
			return &auth.TokenValidityResponse{Ok: true}, nil
		} else {
			return &auth.TokenValidityResponse{Ok: false}, nil
		}
	case <- err := errorChan:

		return &auth.TokenValidityResponse{Ok: false}, err

	case <- ctx.Done():
		return &auth.TokenValidityResponse{Ok: false}, ctx.Err()
	}
	
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	response, _ := DAL.GetUser(ctx,s.kv, req.Username)
	if utilities.CheckPasswordHash(req.Password, response.Password) {
		return &auth.LoginResponse{Ok: true}, nil
	} else {
		return &auth.LoginResponse{Ok: false}, nil
	}

}

func (s *AuthServer) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {

	result := make(chan string)
	errorChan := make(chan error)

	go DAL.DeleteAccessToken(ctx,s.kv, req.Username,result,errorChan)

	select{
	case <- result:
		 return &auth.LogoutResponse{}, nil
	case <- err := errorChan:
		return &auth.LogoutResponse{}, err
	case <- ctx.Done():
		return &auth.LogoutResponse{}, ctx.Err()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3004")
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
	auth.RegisterAuthServiceServer(server, &AuthServer{kv: keyVal})
	reflection.Register(server)
	log.Println("Auth service running on :3004")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
