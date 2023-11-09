package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	commomTools "github.com/yangbiny/reason-commons/common"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/repo"
	"reason-im/pkg/model"
)

type MsgService struct {
	friendDao repo.FriendDao
}

func NewMsgService() *MsgService {
	return &MsgService{}
}

func (msgService *MsgService) SendMsg(c *gin.Context, cmd *MsgCmd) (bool, *apierror.ApiError) {
	msgType := model.CheckIsMsgType(*cmd.ObjectType)
	if !msgType {
		return false, apierror.WhenParamError(fmt.Errorf("invalid object type"))
	}
	if commomTools.IsBlankStr(cmd.Msg) {
		return false, apierror.WhenParamError(fmt.Errorf("消息不能为空（包括纯空字符串）"))
	}

	switch model.MsgType(*cmd.ObjectType) {
	case model.MsgTypeFriend:
		return msgService.sendMsgToFriend(c, cmd)
	case model.MsgTypeGroup:
		return msgService.sendMsgToGroup(c, cmd)
	default:
		return false, apierror.WhenParamError(fmt.Errorf("invalid object type"))
	}
}

func (msgService *MsgService) sendMsgToGroup(c *gin.Context, cmd *MsgCmd) (bool, *apierror.ApiError) {
	panic("implement me")
}

func (msgService *MsgService) sendMsgToFriend(c *gin.Context, cmd *MsgCmd) (bool, *apierror.ApiError) {
	friendId := cmd.ObjectId
	friendInfo, err := msgService.friendDao.QueryFriendInfo(c, cmd.UserId, friendId)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}
	if friendInfo == nil {
		return false, apierror.WhenParamError(fmt.Errorf("他还不是你的朋友"))
	}
	SendMsg(cmd.UserId, cmd.ObjectId, cmd.Msg)
	return true, nil
}

type MsgCmd struct {
	UserId     *int64  `login_user_id:"user_id" validate:"required"`
	ObjectId   *int64  `json:"object_id" validate:"required"`
	ObjectType *int32  `json:"object_type" validate:"required"`
	Msg        *string `json:"msg" validate:"required"`
}
