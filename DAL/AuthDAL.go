package DAL

import (
	// "ds-project/common/proto/models"
	"encoding/json"
	"ds-project/raft"
	"github.com/coreos/etcd/clientv3"
	"context"
	// "fmt"
	"sync"
)

var(
	mutex     sync.Mutex
)

type TokenDB struct {
	Tokens        map[string]string
}

func SetAccessToken(ctx context.Context, kv clientv3.KV, username string, token string,result chan bool, errorChan chan error) {

	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"tokens")
	var result TokenDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        error <- err
        return
    }

    result.Tokens[username] = token
	marshalledToken, err := json.Marshal(result)
	raft.PutKey(ctx,kv,"tokens",marshalledToken)
	result <- true
	return
}

func GetAccessToken(ctx context.Context, kv clientv3.KV, username string, result chan string, errorChan chan error) {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"tokens")
	var result TokenDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        error <- err
        return
    }

	token, _ := result.Tokens[username]
	result <- token
	return
}

func DeleteAccessToken(ctx context.Context, kv clientv3.KV, username string,result chan bool, errorChan chan error) {
	
	mutex.Lock()
	defer mutex.Unlock()
	
	bt := raft.GetKey(ctx,kv,"tokens")
	var result TokenDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        error <- err
        return
    }

   	delete(result.Tokens, username)

	marshalledToken, err := json.Marshal(result)
	raft.PutKey(ctx,kv,"tokens",marshalledToken)
	result <- true
	return
}
