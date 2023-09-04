package rpcclient

import (
	"reason-im/pkg/model"
)

type FriendInvite model.FriendInvite

type FriendInviteDao interface {
	NewFriend(friend *FriendInvite) *FriendInvite
	GetFriendInfo(friendId int64) *FriendInvite
	// UpdateInvite 修改邀请状态
	UpdateInvite(cmd *FriendInvite) bool
	// QueryInviteFriendList 查询用户的 邀请列表
	QueryInviteFriendList(userId int64) []*FriendInvite
}
