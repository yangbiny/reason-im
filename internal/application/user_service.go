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

func (userService UserService) NewUser(user *rpcclient.User) *rpcclient.User {
	return userService.UserDao.NewUser(user)
}

func (userService UserService) GetUserInfo(userId int64) *rpcclient.User {
	return userService.UserDao.GetUserInfo(userId)
}
