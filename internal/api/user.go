package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/pkg/rpcclient"
)

type UserApi rpcclient.UserClientHandler

func (api *UserApi) registerNewUser(c *gin.Context) {

}
