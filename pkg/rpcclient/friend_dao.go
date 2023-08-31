package rpcclient

import "reason-im/internal/entity"

type Friend entity.Friend

type ChangeFriendStatusCmd struct {
	Id     int64
	Status entity.FriendStatus
}

type FriendClient interface {
	NewFriend(friend *Friend) *Friend
	GetFriendInfo(friendId int64) *Friend
	// ChangeFriendStatus 修改朋友状态，包括接收邀请、拒绝邀请和删除等
	ChangeFriendStatus(cmd ChangeFriendStatusCmd) bool
	// QueryInviteFriendList 查询用户的 邀请列表
	QueryInviteFriendList(userId int64) []*Friend
}
