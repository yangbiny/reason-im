package service

var channelMap map[string]*channel

type channel struct {
	ChannelId string
}

func SubChannel(channel string, userId int64) {
	c := channelMap[channel]
	if c == nil {
		return
	}
}
