package utilities

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"
)

func GetKey(ctx context.Context, client *clientv3.Client, key string) []byte {
	response, _ := client.Get(ctx, key)
	var bytes []byte
	log.Println("In get key")
	log.Println(response)
	if response != nil{
	if response.Kvs != nil {
		bytes = []byte(string(response.Kvs[0].Value))
	}
	}
	return bytes
}

func PutKey(ctx context.Context, client *clientv3.Client, key string, value []byte) int64 {
	response, _ := client.Put(ctx, key, string(value))
	rev := response.Header.Revision
	return rev
}
