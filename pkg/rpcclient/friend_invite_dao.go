package rpcclient

import (
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

type FriendInvite model.FriendInvite

type FriendInviteDao interface {
	NewFriend(friend *model.FriendInvite) *FriendInvite
	GetFriendInviteInfo(userId int64, friendId int64) *FriendInvite
	// UpdateInvite 修改邀请状态
	UpdateInvite(cmd *FriendInvite) bool
	// QueryInviteFriendList 查询用户的 邀请列表
	QueryInviteFriendList(userId int64) []*FriendInvite
}

type FriendInviteDaoImpl struct {
	DatabaseTpl *mysql.DatabaseTpl
}

func NewFriendInviteDao(tpl *mysql.DatabaseTpl) FriendInviteDao {
	return FriendInviteDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (f FriendInviteDaoImpl) NewFriend(friend *model.FriendInvite) *FriendInvite {
	//TODO implement me
	panic("implement me")
}

func (f FriendInviteDaoImpl) GetFriendInviteInfo(userId int64, friendId int64) *FriendInvite {
	//TODO implement me
	panic("implement me")
}

func (f FriendInviteDaoImpl) UpdateInvite(cmd *FriendInvite) bool {
	//TODO implement me
	panic("implement me")
}

func (f FriendInviteDaoImpl) QueryInviteFriendList(userId int64) []*FriendInvite {
	//TODO implement me
	panic("implement me")
}
