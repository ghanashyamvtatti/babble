package userdal

import (
	"ds-project/common"
	"ds-project/common/proto/models"
)

type UserDAL interface {
	GetUser(request common.DALRequest, username string, res chan *models.User)
	GetUsers(request common.DALRequest, res chan map[string]*models.User)
	CreateUser(request common.DALRequest, username string, value *models.User, res chan bool)
}
