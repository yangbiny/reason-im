package service

var channelMap map[string]*channel

type channel struct {
	ChannelId string
}

// ApplySubToken 申请订阅的 token
func ApplySubToken(channel string, userId int64) string {
	return ""
}

func SubChannel(channel string, userId int64, token string) {
	c := channelMap[channel]
	if c == nil {
		return
	}
	// 验证 token 是否合法
}

func UnSubChannel(channel string, userId int64) {

}

// SendMsgToChannel 发送消息到频道。订阅者 可以在消息 发送后，接受到发送的消息
func SendMsgToChannel(channel string, msg string) {

}
