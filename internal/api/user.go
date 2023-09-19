package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/rpcclient"
	"reason-im/pkg/service"
	"strconv"
)

type UserApi struct {
	userService *service.UserService
}

func NewUserApi(service *service.UserService) UserApi {
	return UserApi{
		userService: service,
	}
}

func (api *UserApi) RegisterNewUser(c *gin.Context) {
	caller.Call(api.userService.NewUser, c, new(rpcclient.User))
}

func (api *UserApi) QueryUserById(c *gin.Context) {
	var userIdStr, has = c.GetQuery("id")
	if !has {
		caller.ResponseWithParamInvalid(c, "参数不正确")
	}
	userId, _ := strconv.Atoi(userIdStr)
	caller.CallWithParam(api.userService.GetUserInfo, c, int64(userId))
}
