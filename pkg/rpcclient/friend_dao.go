package rpcclient

import (
	"context"
	"fmt"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

var friendTableName = "im_friend"
var friendColumns = "id,user_id,friend_id,status,remark,gmt_create,gmt_update"

type Friend model.Friend

type DeleteFriendCmd struct {
	Id     int64
	Status model.FriendStatus
}

type FriendDao interface {
	NewFriend(friend *Friend) *Friend
	GetFriendInfo(friendId int64) *Friend
	DeleteFriend(cmd DeleteFriendCmd) bool
	QueryFriendList(userId int64) []*Friend
	QueryFriendInfo(userId int64, friendId int64) *Friend
}

type FriendDaoImpl struct {
	DatabaseTpl *mysql2.DatabaseTpl
}

func NewFriendDao(tpl *mysql2.DatabaseTpl) FriendDao {
	return &FriendDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (dao *FriendDaoImpl) NewFriend(friend *Friend) *Friend {
	panic("implement me")
}

func (dao *FriendDaoImpl) GetFriendInfo(friendId int64) *Friend {
	panic("implement me")
}

func (dao *FriendDaoImpl) DeleteFriend(cmd DeleteFriendCmd) bool {
	//TODO implement me
	panic("implement me")
}

func (dao *FriendDaoImpl) QueryFriendList(userId int64) []*Friend {
	//TODO implement me
	panic("implement me")
}

func (dao *FriendDaoImpl) QueryFriendInfo(userId int64, friendId int64) *Friend {
	var sql = fmt.Sprintf("select %s from %s where user_id = ? and friend_id = ?", friendColumns, friendTableName)
	friend := dao.DatabaseTpl.FindOne(context.Background(), sql, Friend{}, userId, friendId).(Friend)
	return &friend
}
