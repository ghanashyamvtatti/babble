package DAL

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

func GetUser(ctx context.Context, kv clientv3.KV, username string,result chan *models.User, errorChan chan error) (*models.User, bool) {

	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"users")
	var r UsersDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }
   
	result <- r.Users[username]
	return
}

func GetUsers(ctx context.Context, kv clientv3.KV,result chan map[string]*models.User, errorChan chan error) map[string]*models.User {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"users")
	var r UsersDB
    err:= json.Unmarshal(bt, &r)
    if err != nil {
        errorChan <- err
        return
    }
	result <- r.Users
	return
}

func CreateUser(ctx context.Context, kv clientv3.KV, username string, result chan bool, errorChan chan error) bool {
	
	mutex.Lock()
	defer mutex.Unlock()

	bt := raft.GetKey(ctx,kv,"users")
	var result UsersDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        errorChan <- err
        return
    }
    fmt.Println("result")
    fmt.Println(result)
	result.Users[username] = value
	marshalledUser, err := json.Marshal(result)
	raft.PutKey(ctx,kv,"users",marshalledUser)
	result <- true 
	return
}
