package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/application"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/rpcclient"
	"strconv"
)

type UserService application.UserService

func NewUserApi(service *application.UserService) UserService {
	return UserService{
		UserDao: service.UserDao,
	}
}

func (api *UserService) RegisterNewUser(c *gin.Context) {
	var user rpcclient.User
	caller.Call(api.UserDao.NewUser, c, &user)
}

func (api *UserService) QueryUserById(c *gin.Context) {
	var userIdStr, has = c.GetQuery("id")
	if !has {
		caller.ResponseWithParamInvalid(c, "参数不正确")
	}
	userId, _ := strconv.Atoi(userIdStr)
	caller.CallWithParam(api.UserDao.GetUserInfo, c, int64(userId))
}
