package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
func (subService *SubService) ApplySubToken(channel string, userId int64) (*string, error) {
	// 可以在这里验证用户是否有权限订阅
	claims := map[string]string{}
	claims["channel"] = claims[channel]
	applyToken, err := token.ApplyToken(context.Background(), userId, claims)
	if err != nil {
		return nil, err
	}
	return applyToken, nil
}

func (subService *SubService) SubChannel(channel string, userId int64, tokenStr *string) error {
	// 验证 token 是否合法
	parseToken, err := token.ParseToken(context.Background(), *tokenStr)
	if err != nil {
		return err
	}
	if parseToken.KeyAsString("channel") != channel {
		return fmt.Errorf("token is invalid")
	}
	if parseToken.UserId() != userId {
		return fmt.Errorf("token is invalid")
	}
	channels := *channelMap[channel]
	subChannels := append(channels, subChannel{ChannelId: channel, UserId: userId})
	channelMap[channel] = &subChannels
	return nil
}

// UnSubChannel 取消订阅
func (subService *SubService) UnSubChannel(channel string, userId int64) {
	channels := *channelMap[channel]
	for i, subChannel := range channels {
		if subChannel.UserId == userId {
			channels = append(channels[:i], channels[i+1:]...)
		}
	}
	channelMap[channel] = &channels
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
