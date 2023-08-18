package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/rpcclient"
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
