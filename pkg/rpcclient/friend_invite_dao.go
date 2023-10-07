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
	NewFriend(ctx context.Context, friend *model.FriendInvite) (*FriendInvite, error)
	GetFriendInviteInfo(ctx context.Context, userId int64, friendId int64) (*FriendInvite, error)
	// UpdateInvite 修改邀请状态
	UpdateInvite(ctx context.Context, cmd *FriendInvite) (bool, error)
	// QueryInviteFriendList 查询用户的 邀请列表
	QueryInviteFriendList(ctx context.Context, userId int64) ([]*FriendInvite, error)
	QueryInvite(ctx context.Context, id int64) (*FriendInvite, error)
}

type FriendInviteDaoImpl struct {
	DatabaseTpl *mysql.DatabaseTpl
}

func NewFriendInviteDao(tpl *mysql.DatabaseTpl) FriendInviteDao {
	return FriendInviteDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (f FriendInviteDaoImpl) QueryInvite(ctx context.Context, id int64) (*FriendInvite, error) {
	var sql = fmt.Sprintf("select %s from %s where id = ?", friendInviteColumns, friendInviteTableName)
	one, err := f.DatabaseTpl.FindOne(ctx, sql, FriendInvite{}, id)
	if one == nil || err != nil {
		return nil, err
	}
	friendInvite := one.(FriendInvite)
	return &friendInvite, nil
}

func (f FriendInviteDaoImpl) NewFriend(ctx context.Context, friend *FriendInvite) (*FriendInvite, error) {
	var sql = fmt.Sprintf("insert into %s (user_id,friend_id,extra,status,gmt_create,gmt_update) values (?,?,?,?,?,?)", friendInviteTableName)
	id, err := f.DatabaseTpl.Insert(ctx, sql, friend.UserId, friend.FriendId, friend.Extra, friend.Status, friend.GmtCreate, friend.GmtUpdate)
	if err != nil {
		return nil, err
	}
	friend.Id = id
	return friend, nil
}

func (f FriendInviteDaoImpl) GetFriendInviteInfo(ctx context.Context, userId int64, friendId int64) (*FriendInvite, error) {
	var sql = fmt.Sprintf("select %s from %s where user_id = ? and friend_id = ?", friendInviteColumns, friendInviteTableName)
	one, err := f.DatabaseTpl.FindOne(ctx, sql, FriendInvite{}, userId, friendId)
	if one == nil || err != nil {
		return nil, err
	}
	friendInvite := one.(FriendInvite)
	return &friendInvite, nil
}

func (f FriendInviteDaoImpl) UpdateInvite(ctx context.Context, cmd *FriendInvite) (bool, error) {
	var sql = fmt.Sprintf("update %s set status = ?,extra = ?,gmt_update where id = ?", friendInviteTableName)
	update, err := f.DatabaseTpl.Update(ctx, sql, cmd.Status, cmd.Extra, cmd.GmtUpdate, cmd.Id)
	if err != nil {
		return false, err
	}
	return update > 0, nil
}

func (f FriendInviteDaoImpl) QueryInviteFriendList(ctx context.Context, userId int64) ([]*FriendInvite, error) {
	var sql = fmt.Sprintf("select %s from %s where user_id = ?", friendInviteColumns, friendInviteTableName)
	one, err := f.DatabaseTpl.FindList(ctx, sql, FriendInvite{}, userId)
	if err != nil {
		return nil, err
	}
	if len(one) == 0 {
		// 没有数据，返回空集合
		return []*FriendInvite{}, nil
	}
	var friendInvites []*FriendInvite
	for _, item := range one {
		friendInvites = append(friendInvites, item.(*FriendInvite))
	}
	return friendInvites, nil
}
