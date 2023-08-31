package application

import (
	"reason-im/pkg/rpcclient"
)

type UserService struct {
	UserDao *rpcclient.UserDaoImpl
}

func NewUserService(userDao *rpcclient.UserDaoImpl) UserService {
	return UserService{
		UserDao: userDao,
	}
}
