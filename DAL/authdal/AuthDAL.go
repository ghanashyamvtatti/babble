package authdal

import (
	"encoding/json"
	"ds-project/raft"
	"github.com/coreos/etcd/clientv3"
	"context"
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
	var r TokenDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }

    r.Tokens[username] = token
	marshalledToken, err := json.Marshal(r)
	raft.PutKey(ctx,kv,"tokens",marshalledToken)
	result <- true
	return
}

func GetAccessToken(ctx context.Context, kv clientv3.KV, username string, result chan string, errorChan chan error) {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"tokens")
	var r TokenDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }

	token, _ := r.Tokens[username]
	result <- token
	return
}

func DeleteAccessToken(ctx context.Context, kv clientv3.KV, username string,result chan bool, errorChan chan error) {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"tokens")
	var r TokenDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }

   	delete(r.Tokens, username)

	marshalledToken, err := json.Marshal(r)
	raft.PutKey(ctx,kv,"tokens",marshalledToken)
	result <- true
	return
}
