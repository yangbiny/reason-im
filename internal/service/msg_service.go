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
	SendMsg(cmd.UserId, cmd.ObjectId, cmd.Msg)
	return true, nil
}

type MsgCmd struct {
	UserId     *int64  `login_user_id:"user_id" validate:"required"`
	ObjectId   *int64  `json:"object_id" validate:"required"`
	ObjectType *int64  `json:"object_type" validate:"required"`
	Msg        *string `json:"msg" validate:"required"`
}
