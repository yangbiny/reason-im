package service

import (
	"github.com/gin-gonic/gin"
	apierror "github.com/yangbiny/reason-commons/err"
)

type MsgService struct {
}

func NewMsgService() *MsgService {
	return &MsgService{}
}

func (msg *MsgService) SendMsg(c *gin.Context, cmd *MsgCmd) (bool, *apierror.ApiError) {
	SendMsg(cmd.ObjectId, cmd.Msg)
	return true, nil
}

type MsgCmd struct {
	UserId   int64  `login_user_id:"user_id" required:"true"`
	ObjectId int64  `json:"object_id" required:"true"`
	Msg      string `json:"msg" required:"true"`
}
