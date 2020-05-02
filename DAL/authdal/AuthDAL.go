package authdal

import (
	"ds-project/common/utilities"
	"ds-project/config"
	"encoding/json"
	"sync"
)

var (
	mutex sync.Mutex
)

type TokenDB struct {
	Tokens map[string]string
}

func SetAccessToken(request config.DALRequest, username string, token string, result chan bool) {
	mutex.Lock()
	defer mutex.Unlock()

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

func GetAccessToken(request config.DALRequest, username string, result chan string) {
	mutex.Lock()
	defer mutex.Unlock()

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

func DeleteAccessToken(request config.DALRequest, username string, result chan bool) {
	mutex.Lock()
	defer mutex.Unlock()

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
