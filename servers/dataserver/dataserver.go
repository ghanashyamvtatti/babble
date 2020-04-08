package dataserver

import (
	"context"
	"ds-project/services"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
	"ds-project/common/proto"
)

type server struct {
	proto.UnimplementedDatagServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":3003")
	if err != nil {
		panic(err)
	}


	// ctx := context.WithValue(context.Backgqround(), "appConfig", appConfig)
	appConfig := config.NewAppConfig()
	users := services.GetAllUsers(appConfig)
	log.Printf(string(len(users)))
	srv := grpc.NewServer()
	proto.RegisterPostServiceServer()
	reflection.Register(srv)
	log.Printf("User service running on :3001")

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

// Define all the DAL methods we might ever require (What the services need)

//