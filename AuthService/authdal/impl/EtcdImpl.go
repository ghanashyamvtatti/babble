package impl

import (
	"ds-project/common"
	"ds-project/common/utilities"
	"encoding/json"
	"sync"
)

type TokenDB struct {
	Tokens map[string]string
}

type EtcdAuthDAL struct {
	Mutex sync.Mutex
}

func (authDAL *EtcdAuthDAL) SetAccessToken(request common.DALRequest, username string, token string, result chan bool) {
	authDAL.Mutex.Lock()
	defer authDAL.Mutex.Unlock()

	bt := utilities.GetKey(request.Ctx, request.Client, "tokens")
	var r TokenDB
	err := json.Unmarshal(bt, &r)
	if err != nil {
		request.ErrorChan <- err
		return
	}

	r.Tokens[username] = token
	marshalledToken, err := json.Marshal(r)
	utilities.PutKey(request.Ctx, request.Client, "tokens", marshalledToken)
	result <- true
	return
}

func (authDAL *EtcdAuthDAL) GetAccessToken(request common.DALRequest, username string, result chan string) {
	authDAL.Mutex.Lock()
	defer authDAL.Mutex.Unlock()

	bt := utilities.GetKey(request.Ctx, request.Client, "tokens")
	var r TokenDB
	err := json.Unmarshal(bt, &r)
	if err != nil {
		request.ErrorChan <- err
		return
	}

	token, _ := r.Tokens[username]
	result <- token
	return
}

func (authDAL *EtcdAuthDAL) DeleteAccessToken(request common.DALRequest, username string, result chan bool) {
	authDAL.Mutex.Lock()
	defer authDAL.Mutex.Unlock()

	bt := utilities.GetKey(request.Ctx, request.Client, "tokens")
	var r TokenDB
	err := json.Unmarshal(bt, &r)
	if err != nil {
		request.ErrorChan <- err
		return
	}

	delete(r.Tokens, username)

	marshalledToken, err := json.Marshal(r)
	utilities.PutKey(request.Ctx, request.Client, "tokens", marshalledToken)
	result <- true
	return
}
