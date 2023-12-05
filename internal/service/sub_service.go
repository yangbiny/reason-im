package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/utils/token"
	"reason-im/pkg/model"
)

type SubService struct {
	msgService MsgService
}

func NewSubService(msgService MsgService) SubService {
	return SubService{msgService: msgService}
}

var channelMap map[string]*[]subChannel

func init() {
	channelMap = make(map[string]*[]subChannel)
}

const sendUserId int64 = 2

type subChannel struct {
	ChannelId string
	UserId    int64
}

// ApplySubToken 申请订阅的 token
func (subService *SubService) ApplySubToken(ctx *gin.Context, req *ApplyTokenReq) (*string, *apierror.ApiError) {
	// 可以在这里验证用户是否有权限订阅
	claims := map[string]string{}
	claims["channel"] = claims[req.Channel]
	applyToken, err := token.ApplyToken(context.Background(), req.UserId, claims)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	return applyToken, nil
}

func (subService *SubService) SubChannel(ctx *gin.Context, req *SubChannelReq) (bool, *apierror.ApiError) {
	// 验证 token 是否合法
	parseToken, err := token.ParseToken(context.Background(), *req.Token)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}
	channel := *req.Channel
	if parseToken.KeyAsString("channel") != channel {
		return false, apierror.WhenParamError(fmt.Errorf("token is invalid"))
	}
	if parseToken.UserId() != req.UserId {
		return false, apierror.WhenParamError(fmt.Errorf("token is invalid"))
	}
	channels := *channelMap[channel]
	subChannels := append(channels, subChannel{ChannelId: channel, UserId: req.UserId})
	channelMap[channel] = &subChannels
	return true, nil
}

// UnSubChannel 取消订阅
func (subService *SubService) UnSubChannel(ctx *gin.Context, req *UnSubChannelReq) (bool, *apierror.ApiError) {
	channel := *req.Channel
	channels := *channelMap[channel]
	for i, subChannel := range channels {
		if subChannel.UserId == req.UserId {
			channels = append(channels[:i], channels[i+1:]...)
		}
	}
	channelMap[channel] = &channels

	return true, nil
}

// SendMsgToChannel 发送消息到频道。订阅者 可以在消息 发送后，接受到发送的消息
func (subService *SubService) SendMsgToChannel(ctx *gin.Context, channel string, msg string) {
	channels := *channelMap[channel]
	var senderUser = sendUserId
	for _, item := range channels {
		SendMsg(&item.UserId, &Msg{
			MsgType:    model.MsgTypeSub,
			ToId:       &item.UserId,
			Msg:        &msg,
			FromUserId: &senderUser,
		})
	}
}

type ApplyTokenReq struct {
	Channel string `json:"channel"`
	UserId  int64  `login_user_id:"user_id"`
}

type SubChannelReq struct {
	Channel *string `json:"channel"`
	UserId  int64   `login_user_id:"user_id"`
	Token   *string `json:"token"`
}

type UnSubChannelReq struct {
	Channel *string `json:"channel"`
	UserId  int64   `login_user_id:"user_id"`
	Token   *string `json:"token"`
}
