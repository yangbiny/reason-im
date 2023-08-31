package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/application"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/rpcclient"
	"strconv"
)

type UserApi struct {
	userService *application.UserService
}

func NewUserApi(service *application.UserService) UserApi {
	return UserApi{
		userService: service,
	}
}

func (api *UserApi) RegisterNewUser(c *gin.Context) {
	var user rpcclient.User
	caller.Call(api.userService.NewUser, c, &user)
}

func (api *UserApi) QueryUserById(c *gin.Context) {
	var userIdStr, has = c.GetQuery("id")
	if !has {
		caller.ResponseWithParamInvalid(c, "参数不正确")
	}
	userId, _ := strconv.Atoi(userIdStr)
	caller.CallWithParam(api.userService.GetUserInfo, c, int64(userId))
}
