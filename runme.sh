#!/bin/sh

# Run etcd nodes in the background

nohup etcd --name s1 \
  --data-dir /tmp/etcd/s1 \
  --listen-client-urls http://localhost:2379 \
  --advertise-client-urls http://localhost:2379 \
  --listen-peer-urls http://localhost:2380 \
  --initial-advertise-peer-urls http://localhost:2380 \
  --initial-cluster s1=http://localhost:2380,s2=http://localhost:22380,s3=http://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new &

nohup etcd --name s2 \
  --data-dir /tmp/etcd/s2 \
  --listen-client-urls http://localhost:22379 \
  --advertise-client-urls http://localhost:22379 \
  --listen-peer-urls http://localhost:22380 \
  --initial-advertise-peer-urls http://localhost:22380 \
  --initial-cluster s1=http://localhost:2380,s2=http://localhost:22380,s3=http://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new &

nohup etcd --name s3 \
  --data-dir /tmp/etcd/s3 \
  --listen-client-urls http://localhost:32379 \
  --advertise-client-urls http://localhost:32379 \
  --listen-peer-urls http://localhost:32380 \
  --initial-advertise-peer-urls http://localhost:32380 \
  --initial-cluster s1=http://localhost:2380,s2=http://localhost:22380,s3=http://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new &

nohup go run servers/dataserver/dataserver.go &
nohup go run AuthService/authserver.go &
nohup go run PostService/postsserver.go &
nohup go run UserService/userserver.go &
nohup go run SubscriptionService/subscriptionsserver.go &
nohup npm start --prefix UI/babble &
go run web.go