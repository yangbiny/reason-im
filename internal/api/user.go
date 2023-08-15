package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils"
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
	utils.Call(api.Client.NewUser, c, &user)
}
