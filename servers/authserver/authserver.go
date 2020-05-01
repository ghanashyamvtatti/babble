package main

import (
	"context"
	"ds-project/common/proto/auth"
	//"ds-project/common/proto/dsl"
	"ds-project/common/utilities"
	"ds-project/DAL"
	// "ds-project/common/proto/users"
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

	DAL.SetAccessToken(ctx,s.kv, req.Username, token.String())


	// _, rpcErr := s.dslClient.SetAccessToken(ctx, &dsl.SetAccessTokenRequest{
	// 	Username: req.Username,
	// 	Token:    token.String(),
	// })
	return &auth.GenerateTokenResponse{Token: token.String()}, nil
}

func (s *AuthServer) CheckAccessTokenValid(ctx context.Context, req *auth.TokenValidityRequest) (*auth.TokenValidityResponse, error) {
	token, _ := DAL.GetAccessToken(ctx,s.kv, req.Username)
	
	if token == req.Token {
		return &auth.TokenValidityResponse{Ok: true}, nil
	} else {
		return &auth.TokenValidityResponse{Ok: false}, nil
	}
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	response, _ := DAL.GetUser(ctx,s.kv, req.Username)
	// response, err := s.dslClient.GetUser(ctx, &dsl.GetUserRequest{Username: req.Username})
	// if err != nil {
	// 	return &auth.LoginResponse{Ok: false}, err
	// } else {
	if utilities.CheckPasswordHash(req.Password, response.Password) {
		return &auth.LoginResponse{Ok: true}, nil
	} else {
		return &auth.LoginResponse{Ok: false}, nil
	}

	// }
}

func (s *AuthServer) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	DAL.DeleteAccessToken(ctx,s.kv, req.Username)
	// if err != nil {
	// 	return &auth.LogoutResponse{}, err
	// } else {
		return &auth.LogoutResponse{}, nil
	// }
}

func main() {
	listener, err := net.Listen("tcp", ":3004")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// // Set up a connection to the DSL server.
	// conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(), grpc.WithBlock())

	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// dslClient := dsl.NewDataServiceClient(conn)

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
