package rpcclient

import (
	"context"
	"fmt"
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

var (
	friendInviteTableName = "im_friend_invite"
	friendInviteColumns   = "id,user_id,friend_id,extra,status,gmt_create,gmt_update"
)

type FriendInvite = model.FriendInvite

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

func (f FriendInviteDaoImpl) NewFriend(friend *FriendInvite) *FriendInvite {
	var sql = fmt.Sprintf("insert into %s (user_id,friend_id,extra,status,gmt_create,gmt_update) values (%s,%s,%s,%s,%s,%s)",
		friendInviteTableName, friend.UserId, friend.FriendId, friend.Extra, friend.Status, friend.GmtCreate, friend.GmtUpdate)
	id := f.DatabaseTpl.Insert(context.Background(), sql)
	friend.Id = id
	return friend
}

func (f FriendInviteDaoImpl) GetFriendInviteInfo(userId int64, friendId int64) *FriendInvite {
	var sql = fmt.Sprintf("select %s from %s where user_id = ? and friend_id = ?", friendInviteColumns, friendInviteTableName)
	one := f.DatabaseTpl.FindOne(context.Background(), sql, FriendInvite{}, userId, friendId)
	if one == nil {
		return nil
	}
	friendInvite := one.(FriendInvite)
	return &friendInvite
}

func (f FriendInviteDaoImpl) UpdateInvite(cmd *FriendInvite) bool {
	//TODO implement me
	panic("implement me")
}

func (f FriendInviteDaoImpl) QueryInviteFriendList(userId int64) []*FriendInvite {
	//TODO implement me
	panic("implement me")
}
