package main

import (  
    "fmt"
    "context"
    "ds-project/sandbox"
    "ds-project/common/proto/models"
    "encoding/json"
    "github.com/coreos/etcd/clientv3"
    "time"
    // "github.com/golang/protobuf/ptypes"
    // "reflect"
)

var (  
    dialTimeout    = 2 * time.Second
    requestTimeout = 10 * time.Second
)

type UsersDB struct {
	Users         map[string]*models.User
}


func main() {  
    ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
    cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    kv := clientv3.NewKV(cli)


	// // var result UsersDB

	// result := &UsersDB{
	// 	Users:         map[string]*models.User{},
	// }

 //    result.Users["ghanu"] =  &models.User{
	// 	FullName:  "Ghanashyam",
	// 	Password:  "$2a$14$YJHc.LklumtVpMb1wl6GweagO/4WqwXFOMylc4oOFP/iufqVwMOAK",
	// 	CreatedAt: ptypes.TimestampNow(),
	// 	UpdatedAt: ptypes.TimestampNow(),
	// }
	// marshalledUser, err := json.Marshal(result)
	// if err != nil {
 //        panic(err)
 //    }
	// RAFT.PutKey(ctx,kv,"users",marshalledUser)

	bt := RAFT.GetKey(ctx,kv,"users")
	var c UsersDB
    er:= json.Unmarshal(bt, &c)
    if er != nil {
        panic(er)
    }
   
	// user, _ := c.Users["ghanu"]
	fmt.Println(c)


    // bt := getKey(ctx,kv,"key")
    // fmt.Println(bt)
    // var result []Post
    // er:= json.Unmarshal(bt, &result)
    // if er != nil {
    //     panic(er)
    // }
    // fmt.Println(result)
}