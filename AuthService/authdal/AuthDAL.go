package authdal

import "ds-project/common"

type AuthDAL interface {
	SetAccessToken(request common.DALRequest, username string, token string, result chan bool)
	GetAccessToken(request common.DALRequest, username string, result chan string)
	DeleteAccessToken(request common.DALRequest, username string, result chan bool)
}