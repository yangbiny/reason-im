package model

type MsgType int

const (
	// MsgTypeFriend 1. 好友消息
	MsgTypeFriend MsgType = iota
	// MsgTypeGroup 2. 群组消息
	MsgTypeGroup
)

func CheckIsMsgType(msgType int32) bool {
	switch msgType {
	case int32(MsgTypeFriend):
		return true
	case int32(MsgTypeGroup):
		return true
	default:
		return false
	}
}
