package config

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

type DALRequest struct {
	Ctx       context.Context
	Client    *clientv3.Client
	ErrorChan chan error
}
