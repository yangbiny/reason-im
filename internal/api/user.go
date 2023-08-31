package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/rpcclient"
	"strconv"
)

type UserApi rpcclient.UserClientHandler

func NewUserApi() UserApi {
	return UserApi{
		Client: rpcclient.UserClientHandler{},
	}
}

func (api *UserApi) RegisterNewUser(c *gin.Context) {
	var user rpcclient.User
	caller.Call(api.Client.NewUser, c, &user)
}

func (api *UserApi) QueryUserById(c *gin.Context) {
	var userIdStr, has = c.GetQuery("id")
	if !has {
		caller.ResponseWithParamInvalid(c, "参数不正确")
	}
	userId, _ := strconv.Atoi(userIdStr)
	caller.CallWithParam(api.Client.GetUserInfo, c, int64(userId))
}
