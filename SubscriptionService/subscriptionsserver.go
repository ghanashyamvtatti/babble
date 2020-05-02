package main

import (
	"context"
	"ds-project/SubscriptionService/subscriptiondal"
	"ds-project/common"
	"ds-project/common/proto/subscriptions"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type SubscriptionServer struct {
	subscriptions.UnimplementedSubscriptionServiceServer
	client *clientv3.Client
}

var (
	dialTimeout = 2 * time.Second
)

/*
rpc Subscribe(SubscribeRequest) returns (SubscribeResponse);
rpc Unsubscribe(SubscribeRequest) returns (SubscribeResponse);
rpc GetSubscriptions(GetSubscriptionsRequest) returns (GetSubscriptionsResponse);
*/

func (s *SubscriptionServer) Subscribe(ctx context.Context, req *subscriptions.SubscribeRequest) (*subscriptions.SubscribeResponse, error) {
	result := make(chan bool)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    s.client,
		ErrorChan: errorChan,
	}
	go subscriptiondal.Subscribe(request, req.Subscriber, req.Publisher, result)

	select {
	case <-result:
		return &subscriptions.SubscribeResponse{}, nil
	case err := <-errorChan:
		return &subscriptions.SubscribeResponse{}, err
	case <-ctx.Done():
		return &subscriptions.SubscribeResponse{}, ctx.Err()
	}
}

func (s *SubscriptionServer) Unsubscribe(ctx context.Context, req *subscriptions.SubscribeRequest) (*subscriptions.SubscribeResponse, error) {
	result := make(chan bool)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       ctx,
		Client:    s.client,
		ErrorChan: errorChan,
	}

	go subscriptiondal.Unsubscribe(request, req.Subscriber, req.Publisher, result)

	select {
	case <-result:
		return &subscriptions.SubscribeResponse{}, nil
	case err := <-errorChan:
		return &subscriptions.SubscribeResponse{}, err
	case <-ctx.Done():
		return &subscriptions.SubscribeResponse{}, ctx.Err()
	}
}

func (s *SubscriptionServer) GetSubscriptions(ctx context.Context, req *subscriptions.GetSubscriptionsRequest) (*subscriptions.GetSubscriptionsResponse, error) {
	result := make(chan *subscriptions.GetSubscriptionsResponse)
	errorChan := make(chan error)
	request := common.DALRequest{
		Ctx:       ctx,
		Client:    s.client,
		ErrorChan: errorChan,
	}
	go subscriptiondal.GetSubscriptions(request, req.Username, result)

	select {
	case res := <-result:
		return res, nil
	case err := <-errorChan:
		return &subscriptions.GetSubscriptionsResponse{}, err
	case <-ctx.Done():
		return &subscriptions.GetSubscriptionsResponse{}, ctx.Err()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Set up a connection to etcd.
	cli, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()

	server := grpc.NewServer()
	subscriptions.RegisterSubscriptionServiceServer(server, &SubscriptionServer{
		client: cli,
	})
	reflection.Register(server)
	log.Println("Subscription service running on :3005")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
