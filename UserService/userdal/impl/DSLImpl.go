package impl

import (
	"ds-project/UserService/config"
	"ds-project/common"
	"ds-project/common/proto/models"
	"sync"
)

type DSLUserDAL struct {
	Mutex     sync.Mutex
	AppConfig *config.ApplicationConfig
}

func (dal *DSLUserDAL) GetUser(request common.DALRequest, username string, res chan *models.User) {

	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	res <- dal.AppConfig.Users[username]
	return
}

func (dal *DSLUserDAL) GetUsers(request common.DALRequest, res chan map[string]*models.User) {

	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()
	res <- dal.AppConfig.Users
	return
}

func (dal *DSLUserDAL) CreateUser(request common.DALRequest, username string, value *models.User, res chan bool) {

	dal.Mutex.Lock()
	defer dal.Mutex.Unlock()

	dal.AppConfig.Users[username] = value
	res <- true
	return
}
