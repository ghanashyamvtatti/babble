package main

import (
	"context"
	"ds-project/common/proto/dsl"
	"ds-project/common/proto/subscriptions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

type SubscriptionServer struct {
	subscriptions.UnimplementedSubscriptionServiceServer
	dslClient dsl.DataServiceClient
	mutex     sync.Mutex
}

/*
rpc Subscribe(SubscribeRequest) returns (SubscribeResponse);
rpc Unsubscribe(SubscribeRequest) returns (SubscribeResponse);
rpc GetSubscriptions(GetSubscriptionsRequest) returns (GetSubscriptionsResponse);
*/

func (s *SubscriptionServer) Subscribe(ctx context.Context, req *subscriptions.SubscribeRequest) (*subscriptions.SubscribeResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, err := s.dslClient.Subscribe(ctx, &dsl.SubscribeRequest{
		Subscriber: req.Subscriber,
		Publisher:  req.Publisher,
	})
	return &subscriptions.SubscribeResponse{}, err
}

func (s *SubscriptionServer) Unsubscribe(ctx context.Context, req *subscriptions.SubscribeRequest) (*subscriptions.SubscribeResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, err := s.dslClient.Unsubscribe(ctx, &dsl.SubscribeRequest{
		Subscriber: req.Subscriber,
		Publisher:  req.Publisher,
	})
	return &subscriptions.SubscribeResponse{}, err
}

func (s *SubscriptionServer) GetSubscriptions(ctx context.Context, req *subscriptions.GetSubscriptionsRequest) (*subscriptions.GetSubscriptionsResponse, error) {
	response, err := s.dslClient.GetSubscriptions(ctx, &dsl.GetSubscriptionsRequest{Username: req.Username})
	if err != nil {
		return &subscriptions.GetSubscriptionsResponse{}, err
	} else {
		return &subscriptions.GetSubscriptionsResponse{Subscriptions: response.Subscriptions}, err
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3005")
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
	subscriptions.RegisterSubscriptionServiceServer(server, &SubscriptionServer{dslClient: dslClient})
	reflection.Register(server)
	log.Println("Subscription service running on :3005")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
