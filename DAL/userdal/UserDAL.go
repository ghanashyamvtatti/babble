package userdal

import (
	"ds-project/common/proto/models"
	"ds-project/common/utilities"
	"ds-project/config"
	"encoding/json"
	"fmt"
	"sync"
)

type UsersDB struct {
	Users map[string]*models.User
}

var (
	mutex sync.Mutex
)

func GetUser(request config.DALRequest, username string, res chan *models.User) {

	mutex.Lock()
	defer mutex.Unlock()

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

func GetUsers(request config.DALRequest, res chan map[string]*models.User) {

	mutex.Lock()
	defer mutex.Unlock()

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

func CreateUser(request config.DALRequest, username string, value *models.User, res chan bool) {

	mutex.Lock()
	defer mutex.Unlock()

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
