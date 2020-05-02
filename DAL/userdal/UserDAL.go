package userdal

import (
	"ds-project/common/proto/models"
	"encoding/json"
	"ds-project/raft"
	"github.com/coreos/etcd/clientv3"
	"context"
	"fmt"
	"sync"
)

type UsersDB struct {
	Users         map[string]*models.User 
}

var(
	mutex     sync.Mutex
)

func GetUser(ctx context.Context, kv clientv3.KV, username string,res chan *models.User, errorChan chan error)  {

	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"users")
	var r UsersDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }
   
	res <- r.Users[username]
	return
}

func GetUsers(ctx context.Context, kv clientv3.KV,res chan map[string]*models.User, errorChan chan error)  {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"users")
	var r UsersDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }
	res <- r.Users
	return
}

func CreateUser(ctx context.Context, kv clientv3.KV, username string,value *models.User, res chan bool, errorChan chan error)  {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"users")
	var r UsersDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }
    fmt.Println("res")
    fmt.Println(r)
	r.Users[username] = value
	marshalledUser, err := json.Marshal(r)
	raft.PutKey(ctx,kv,"users",marshalledUser)
	res <- true 
	return
}
