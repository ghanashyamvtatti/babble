package raft

import (  
    // "fmt"
    "context"
    // "encoding/json"
    "github.com/coreos/etcd/clientv3"
    // "time"
    // "reflect"
)

// type Post struct {
//     Id      int    `json:"id"`
//     UserId int `json:"user_id" binding:"required"`
//     Post string `json:"name" binding:"required"`
//     CreatedAt  time.Time
//     UpdatedAt  time.Time
// }

// var postList = []Post{
//   Post{Id: 1, UserId: 1, Post: "ABC"},
//   Post{Id: 2, UserId: 1, Post: "DEF"},
//   Post{Id: 3, UserId: 2, Post: "GHI"},
//   Post{Id: 4, UserId: 2, Post: "UVW"},
//   Post{Id: 5, UserId: 3, Post: "XYZ"},
// }

// var (  
//     dialTimeout    = 2 * time.Second
//     requestTimeout = 10 * time.Second
// )

// func main() {  
//     ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
//     cli, _ := clientv3.New(clientv3.Config{
//         DialTimeout: dialTimeout,
//         Endpoints: []string{"127.0.0.1:2379"},
//     })
//     defer cli.Close()
//     kv := clientv3.NewKV(cli)
//     bt := getKey(ctx,kv,"key")
//     fmt.Println(bt)
//     var result []Post
//     er:= json.Unmarshal(bt, &result)
//     if er != nil {
//         panic(er)
//     }
//     fmt.Println(result)
// }


// func GetSingleValueDemo(ctx context.Context, kv clientv3.KV) {  
//     fmt.Println("*** GetSingleValueDemo()")
//     // Delete all keys
//     kv.Delete(ctx, "key", clientv3.WithPrefix())

//     b, err := json.Marshal(postList)
//     fmt.Println("This is the type:")
//     fmt.Println(reflect.TypeOf(b))

//     if err != nil {
//         panic(err)
//     }

//     rev := putKey(ctx,kv,"key",b)
//     fmt.Println("Revision:", rev)
//     fmt.Println(reflect.TypeOf(rev))


//     bt := getKey(ctx,kv,"key")
//     var result []Post
//     er:= json.Unmarshal(bt, &result)
//     if er != nil {
//         panic(err)
//     }
//     fmt.Println(result[0].Post)

// }


func GetKey(ctx context.Context, kv clientv3.KV, key string) []byte {
	gr, _ := kv.Get(ctx, key)
	bt := []byte(string(gr.Kvs[0].Value))
	return bt
}

func PutKey(ctx context.Context, kv clientv3.KV, key string, value []uint8) int64{
	pr, _ := kv.Put(ctx, key, string(value))
    rev := pr.Header.Revision
    return rev
}
