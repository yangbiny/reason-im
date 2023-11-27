package service

import (
	"context"
	"fmt"
	"reason-im/internal/utils/token"
)

var channelMap map[string]*channel

type channel struct {
	ChannelId string
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
	c := channelMap[channel]
	if c == nil {
		// 不需要报错，直接 不订阅就可以了
		return nil
	}
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
	return nil
}

func UnSubChannel(channel string, userId int64) {

}

// SendMsgToChannel 发送消息到频道。订阅者 可以在消息 发送后，接受到发送的消息
func SendMsgToChannel(channel string, msg string) {

}
