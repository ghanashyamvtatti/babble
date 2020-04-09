package main

import (
	"context"
	"ds-project/common/proto/auth"
	"ds-project/common/proto/dsl"
	"ds-project/common/utilities"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	dslClient dsl.DataServiceClient
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
	_, rpcErr := s.dslClient.SetAccessToken(ctx, &dsl.SetAccessTokenRequest{
		Username: req.Username,
		Token:    token.String(),
	})
	return &auth.GenerateTokenResponse{Token: token.String()}, rpcErr
}

func (s *AuthServer) CheckAccessTokenValid(ctx context.Context, req *auth.TokenValidityRequest) (*auth.TokenValidityResponse, error) {
	response, err := s.dslClient.GetAccessToken(ctx, &dsl.AccessTokenRequest{Username: req.Username})
	if err != nil {
		return &auth.TokenValidityResponse{Ok: false}, err
	}
	if response.Token == req.Token {
		return &auth.TokenValidityResponse{Ok: true}, err
	} else {
		return &auth.TokenValidityResponse{Ok: false}, err
	}
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	response, err := s.dslClient.GetUser(ctx, &dsl.GetUserRequest{Username: req.Username})
	if err != nil {
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
	_, err := s.dslClient.DeleteAccessToken(ctx, &dsl.AccessTokenRequest{Username: req.Username})
	if err != nil {
		return &auth.LogoutResponse{}, err
	} else {
		return &auth.LogoutResponse{}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3004")
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
	auth.RegisterAuthServiceServer(server, &AuthServer{dslClient: dslClient})
	reflection.Register(server)
	log.Println("Auth service running on :3004")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
