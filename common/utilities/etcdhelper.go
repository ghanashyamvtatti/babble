package utilities

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

func GetKey(ctx context.Context, client *clientv3.Client, key string) []byte {
	response, _ := client.Get(ctx, key)
	var bytes []byte
	if response != nil && response.Kvs != nil {
		bytes = []byte(string(response.Kvs[0].Value))
	}
	return bytes
}

func PutKey(ctx context.Context, client *clientv3.Client, key string, value []byte) int64 {
	response, _ := client.Put(ctx, key, string(value))
	rev := response.Header.Revision
	return rev
}
