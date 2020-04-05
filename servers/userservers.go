package servers

import (
	"context"
	"log"
	"ds-project/common/proto"
	"ds-project/models"
	"net"
	"ds-project/services"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	userproto.RegisterUserServiceServer(srv, &server{})
	reflection.Register(srv)
	log.Printf("User service running on :3001")

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Register(ctx context.Context, request *userproto.REGISTERUSER) (*userproto.USER, error) {
	log.Printf("Register service")

	username := request.GetUserName()
	fullname := request.GetFullName()
	password := request.GetPassword()

	user := services.CheckUserNameExists(appConfig, username,models.User{FullName: fullname, Password: password})


	dt := &userproto.USER{UserName: string(username),FullName: string(fullname)}
	log.Printf("Register response")
	return &userproto.USER{Data:dt}, nil
}

func (s *server) SignIn(ctx context.Context, request *userproto.LOGIN) (*userproto.VALID, error) {
	log.Printf("Sign in user")
	isLogin := services.Login(appConfig, request.GetUserName(), request.GetPassword() )
	return &userproto.VALID{IsValid: isLogin }, nil
}

func (s *server) CheckUserNameExists(ctx context.Context, request *userproto.USERNAME) (*userproto.VALID, error) {
	log.Printf("CheckUserNameExists in user")
	isLogin := services.CheckUserNameExists(appConfig, request.GetUserName())
	return &userproto.VALID{IsValid: isLogin }, nil
}

func (s *server) GenerateAccessToken(ctx context.Context, request *userproto.USERNAME) (*userproto.TOKEN, error) {
	log.Printf("GenerateAccessToken in user")
	token := services.GenerateAccessToken(appConfig, request.GetUserName())
	return &userproto.TOKEN{Token: token }, nil
}

func (s *server) CheckAccessTokenValid(ctx context.Context, request *userproto.USERNAME) (*userproto.TOKEN, error) {
	log.Printf("CheckAccessTokenValid in user")
	isValid := services.CheckAccessTokenValid(appConfig, request.GetUserName(), request.GetToken())
	return &userproto.VALID{IsValid: isValid }, nil
}

