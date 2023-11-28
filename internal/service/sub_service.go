package service

import (
	"context"
	"fmt"
	"reason-im/internal/utils/token"
)

var channelMap map[string]*[]subChannel

func init() {
	channelMap = make(map[string]*[]subChannel)
}

type subChannel struct {
	ChannelId string
	UserId    int64
}

// ApplySubToken 申请订阅的 token
func ApplySubToken(channel string, userId int64) (*string, error) {
	// 可以在这里验证用户是否有权限订阅
	claims := map[string]string{}
	claims["channel"] = claims[channel]
	applyToken, err := token.ApplyToken(context.Background(), userId, claims)
	if err != nil {
		return nil, err
	}
	return applyToken, nil
}

func SubChannel(channel string, userId int64, tokenStr *string) error {
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
func UnSubChannel(channel string, userId int64) {
	channels := *channelMap[channel]
	for i, subChannel := range channels {
		if subChannel.UserId == userId {
			channels = append(channels[:i], channels[i+1:]...)
		}
	}
	channelMap[channel] = &channels
}

// SendMsgToChannel 发送消息到频道。订阅者 可以在消息 发送后，接受到发送的消息
func SendMsgToChannel(channel string, msg string) {

}
