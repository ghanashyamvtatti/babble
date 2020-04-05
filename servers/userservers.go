package main

import (
	"context"
	"log"
	"ds-project/common/proto"
	"ds-project/models"
	"net"
	"ds-project/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ds-project/config"
)

type server struct{}
var appConfig *config.ApplicationConfig

func main() {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
    
 
    // ctx := context.WithValue(context.Background(), "appConfig", appConfig)
    appConfig := config.NewAppConfig()
    users := services.GetAllUsers(appConfig)
    log.Printf(string(len(users)))
	srv := grpc.NewServer()
	userproto.RegisterUserServiceServer(srv, &server{})
	reflection.Register(srv)
	log.Printf("User service running on :3001")

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Register(ctx context.Context, request *userproto.REGISTER) (*userproto.USER, error) {
	log.Printf("Register service")

	username := request.GetUsername()
	fullname := request.GetFullname()
	password := request.GetPassword()

	services.CreateUser(appConfig, username,models.User{FullName: fullname, Password: password})


	// dt := &userproto.USER{Username: string(username),Fullname: string(fullname)}
	log.Printf("Register response")
	return &userproto.USER{Username: string(username),Fullname: string(fullname)}, nil
}

func (s *server) SignIn(ctx context.Context, request *userproto.LOGIN) (*userproto.VALID, error) {
	log.Printf("Sign in user")
	isLogin := services.Login(appConfig, request.GetUsername(), request.GetPassword() )
	return &userproto.VALID{IsValid: isLogin }, nil
}

func (s *server) CheckUserNameExists(ctx context.Context, request *userproto.USERNAME) (*userproto.VALID, error) {
	log.Printf("CheckUserNameExists in user")
	isLogin := services.CheckUserNameExists(appConfig, request.GetUsername())
	return &userproto.VALID{IsValid: isLogin }, nil
}

func (s *server) GenerateAccessToken(ctx context.Context, request *userproto.USERNAME) (*userproto.TOKEN, error) {
	log.Printf("GenerateAccessToken in user")
	token := services.GenerateAccessToken(appConfig, request.GetUsername())
	return &userproto.TOKEN{Token: token }, nil
}

func (s *server) CheckAccessTokenValid(ctx context.Context, request *userproto.CHECKACCESSTOKEN) (*userproto.VALID, error) {
	log.Printf("CheckAccessTokenValid in user")
	isValid := services.CheckAccessTokenValid(appConfig, request.GetUsername(), request.GetToken())
	return &userproto.VALID{IsValid: isValid }, nil
}


func (s *server) GetUsers(ctx context.Context, request *userproto.REQ) (*userproto.USERS, error) {
	log.Printf("GetUsers in user")
	users := services.GetAllUsers(appConfig)
	log.Printf("after app config")

	var result []*userproto.USER
	for i := 0; i < len(users); i++ {
		dt := &userproto.USER{Username: string(users[i].Username),Fullname: string(users[i].FullName)}
		result = append(result, dt)
	}
	log.Printf("Get Users service response")
	return &userproto.USERS{Users: result }, nil
}

