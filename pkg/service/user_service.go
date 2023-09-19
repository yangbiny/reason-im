package service

import (
	"reason-im/pkg/rpcclient"
)

type UserService struct {
	UserDao rpcclient.UserDao
}

func NewUserService(userDao rpcclient.UserDao) UserService {
	return UserService{
		UserDao: userDao,
	}
}

func (userService UserService) NewUser(user *rpcclient.User) (*rpcclient.User, error) {
	return userService.UserDao.NewUser(user)
}

func (userService UserService) GetUserInfo(userId int64) (*rpcclient.User, error) {
	return userService.UserDao.GetUserInfo(userId)
}
