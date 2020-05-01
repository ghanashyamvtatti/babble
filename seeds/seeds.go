package main

import (  
    "fmt"
    "context"
    "ds-project/sandbox"
    "ds-project/common/proto/models"
    "encoding/json"
    "github.com/coreos/etcd/clientv3"
    "time"
    "github.com/golang/protobuf/ptypes"
    // "reflect"
)

var (  
    dialTimeout    = 2 * time.Second
    requestTimeout = 10 * time.Second
)

type UsersDB struct {
	Users         map[string]*models.User
}

type TokenDB struct {
    Tokens        map[string]string
}


func main() {  
    ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
    cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    kv := clientv3.NewKV(cli)

    kv.Delete(ctx, "users", clientv3.WithPrefix())
    kv.Delete(ctx, "tokens", clientv3.WithPrefix())

	users := &UsersDB{
		Users:         map[string]*models.User{},
	}

    tokens := &TokenDB{
        Tokens:        map[string]string{},
    }

    users.Users["ghanu"] =  &models.User{
		FullName:  "Ghanashyam",
		Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

    users.Users["varun"] = &models.User{
        FullName:  "Varun",
        Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
        CreatedAt: ptypes.TimestampNow(),
        UpdatedAt: ptypes.TimestampNow(),
    }

    // Add User 3
    users.Users["pratik"] = &models.User{
        FullName:  "Pratik",
        Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
        CreatedAt: ptypes.TimestampNow(),
        UpdatedAt: ptypes.TimestampNow(),
    }

	marshalledUser, err := json.Marshal(users)
	if err != nil {
        panic(err)
    }
	RAFT.PutKey(ctx,kv,"users",marshalledUser)

	bt := RAFT.GetKey(ctx,kv,"users")
	var c UsersDB
    er:= json.Unmarshal(bt, &c)
    if er != nil {
        panic(er)
    }
   
	fmt.Println(c)

    tokens.Tokens["ghanu"] = "MASTER-TOKEN"
    tokens.Tokens["varun"] = "MASTER-TOKEN"
    tokens.Tokens["pratik"] = "MASTER-TOKEN"

    marshalledToken, err := json.Marshal(tokens)
    if err != nil {
        panic(err)
    }
    RAFT.PutKey(ctx,kv,"tokens",marshalledToken)

    bt1 := RAFT.GetKey(ctx,kv,"tokens")
    var c1 TokenDB
    er1:= json.Unmarshal(bt1, &c1)
    if er1 != nil {
        panic(er1)
    }
   
    fmt.Println(c1)

}