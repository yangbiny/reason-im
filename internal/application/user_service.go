package application

import (
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
)

type UserService struct {
	UserDao rpcclient.UserDao
}

func NewUserService(tpl *mysql.DatabaseTpl) UserService {
	return UserService{
		UserDao: rpcclient.UserDaoImpl{
			DatabaseTpl: tpl,
		},
	}
}
