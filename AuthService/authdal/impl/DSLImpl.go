package impl

import (
	"ds-project/AuthService/config"
	"ds-project/common"
	"sync"
)

type DSLAuthDAL struct {
	Mutex     sync.Mutex
	AppConfig *config.ApplicationConfig
}

func (authDAL *DSLAuthDAL) SetAccessToken(request common.DALRequest, username string, token string, result chan bool) {
	authDAL.Mutex.Lock()
	defer authDAL.Mutex.Unlock()

	authDAL.AppConfig.Tokens[username] = token
	result <- true
	return
}

func (authDAL *DSLAuthDAL) GetAccessToken(request common.DALRequest, username string, result chan string) {
	authDAL.Mutex.Lock()
	defer authDAL.Mutex.Unlock()

	token, _ := authDAL.AppConfig.Tokens[username]
	result <- token
	return
}

func (authDAL *DSLAuthDAL) DeleteAccessToken(request common.DALRequest, username string, result chan bool) {
	authDAL.Mutex.Lock()
	defer authDAL.Mutex.Unlock()

	if _, ok := authDAL.AppConfig.Tokens[username]; ok {
		delete(authDAL.AppConfig.Tokens, username)
	}
	result <- true
	return
}
