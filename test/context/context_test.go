package context

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"ds-project/UserService/userdal"
	"ds-project/common/proto/models"
	"ds-project/common"
)


var (  
    dialTimeout    = 2 * time.Second
    requestTimeout = 10 * time.Second
)
// Test cases for User Service

func TestUserDALStorageGetUsers(t *testing.T) {
	log.Println("Testing User DAL Storage")
	cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    // keyVal := clientv3.NewKV(cli)

	res := make(chan *models.User)
	errorChan := make(chan error)

	request := common.DALRequest{
		Ctx:       context.Background(),
		Client:    cli,
		ErrorChan: errorChan,
	}

	ctx := context.Background()
	dal := userdal.UserDAL{Mutex: sync.Mutex{}}

	go dal.GetUser(request, "ghanu", res)

	select{
	case us := <-res:
		fmt.Println(us)
	case err := <-errorChan:
		fmt.Println(err)
		t.Error("fails")
	case <- ctx.Done():
		fmt.Println(ctx.Done())
		t.Error("fails")
	}
}

func TestUserDALStorageGetUsersWithCancelledContext(t *testing.T) {
	log.Println("Testing User DAL Storage with Cancelled Context")
	cli, _ := clientv3.New(clientv3.Config{
        DialTimeout: dialTimeout,
        Endpoints: []string{"127.0.0.1:2379"},
    })
    defer cli.Close()
    // keyVal := clientv3.NewKV(cli)

	res := make(chan *models.User)
	errorChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	//ctx, _ := context.WithTimeout(context.Background(),time.Duration(2*time.Second))

    


	request := common.DALRequest{
		Ctx:       ctx,
		Client:    cli,
		ErrorChan: errorChan,
	}

	
	dal := userdal.UserDAL{Mutex: sync.Mutex{}}

	// time.Sleep(3 * time.Second)
	go dal.GetUser(request, "ghanu", res)


	select{
	case us := <-res:
		fmt.Println(us)
	case err := <-errorChan:
		fmt.Println(err)
		t.Error("fails")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("Cancelled context case")
		// fmt.Println(ctx.Done())
		// t.Error("fails")
	}

	//defer cancel()
}

