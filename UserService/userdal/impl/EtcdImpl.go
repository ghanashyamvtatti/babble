package impl

import (
	"ds-project/common"
	"ds-project/common/proto/models"
	"ds-project/common/utilities"
	"encoding/json"
	"fmt"
	"sync"
)

type UsersDB struct {
	Users map[string]*models.User
}

type EtcdUserDAL struct {
	Mutex sync.Mutex
}

func (dal *EtcdUserDAL) GetUser(request common.DALRequest, username string, res chan *models.User) {

	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	bt := utilities.GetKey(request.Ctx, request.Client, "users")
	var r UsersDB
	err := json.Unmarshal(bt, &r)
	if err != nil {
		request.ErrorChan <- err
		return
	}

	res <- r.Users[username]
	return
}

func (dal *EtcdUserDAL) GetUsers(request common.DALRequest, res chan map[string]*models.User) {

	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	bt := utilities.GetKey(request.Ctx, request.Client, "users")
	var r UsersDB
	err := json.Unmarshal(bt, &r)
	if err != nil {
		request.ErrorChan <- err
		return
	}
	res <- r.Users
	return
}

func (dal *EtcdUserDAL) CreateUser(request common.DALRequest, username string, value *models.User, res chan bool) {

	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	bt := utilities.GetKey(request.Ctx, request.Client, "users")
	var r UsersDB
	err := json.Unmarshal(bt, &r)
	if err != nil {
		request.ErrorChan <- err
		return
	}
	fmt.Println("res")
	fmt.Println(r)
	r.Users[username] = value
	marshalledUser, err := json.Marshal(r)
	utilities.PutKey(request.Ctx, request.Client, "users", marshalledUser)
	res <- true
	return
}
