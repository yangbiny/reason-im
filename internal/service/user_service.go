package service

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/config/web"
	"reason-im/internal/repo"
)

type UserService struct {
	UserDao repo.UserDao
}

func NewUserService(userDao repo.UserDao) UserService {
	return UserService{
		UserDao: userDao,
	}
}

func (userService UserService) NewUser(ctx *gin.Context, cmd *NewUserCmd) (*repo.User, error) {
	// 创建用户信息
	user := repo.User{
		Name: cmd.Name,
	}
	return userService.UserDao.NewUser(&user)
}

func (userService UserService) GetUserInfo(ctx *gin.Context, userId *int64) (*repo.User, error) {
	return userService.UserDao.GetUserInfo(*userId)
}

func (userService UserService) Login(c *gin.Context, user *UserLoginCmd) (bool, error) {
	queryUser, err := userService.UserDao.QueryUserByName(user.Name)
	if err != nil {
		return false, err
	}
	if queryUser == nil {
		return false, nil
	}
	err = web.GenerateJwtToken(c, queryUser.Name, queryUser.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

type UserLoginCmd struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type NewUserCmd struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}