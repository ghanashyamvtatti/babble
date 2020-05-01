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

func GetUser(ctx context.Context, kv clientv3.KV, username string) (*models.User, bool) {
	bt := raft.GetKey(ctx,kv,"users")
	var result UsersDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        panic(err)
    }
   
	user, ok := result.Users[username]
	return user, ok
}

func GetUsers(ctx context.Context, kv clientv3.KV) map[string]*models.User {
	bt := raft.GetKey(ctx,kv,"users")
	var result UsersDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        panic(err)
    }
	return result.Users
}

func CreateUser(ctx context.Context, kv clientv3.KV, username string, value *models.User) bool {
	bt := raft.GetKey(ctx,kv,"users")
	var result UsersDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        panic(err)
    }
    fmt.Println("result")
    fmt.Println(result)
	result.Users[username] = value
	marshalledUser, err := json.Marshal(result)
	raft.PutKey(ctx,kv,"users",marshalledUser)
	return true
}
