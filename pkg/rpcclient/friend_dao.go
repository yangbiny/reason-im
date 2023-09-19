package rpcclient

import (
	"context"
	"fmt"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

var friendTableName = "im_friend"
var friendColumns = "id,user_id,friend_id,status,remark,gmt_create,gmt_update"

type Friend = model.Friend

type DeleteFriendCmd struct {
	Id     int64
	Status model.FriendStatus
}

type FriendDao interface {
	NewFriend(friend *Friend) (*Friend, error)
	GetFriendInfo(friendId int64) (*Friend, error)
	DeleteFriend(cmd DeleteFriendCmd) (bool, error)
	QueryFriendList(userId int64) ([]*Friend, error)
	QueryFriendInfo(userId int64, friendId int64) (*Friend, error)
}

type FriendDaoImpl struct {
	DatabaseTpl *mysql2.DatabaseTpl
}

func NewFriendDao(tpl *mysql2.DatabaseTpl) FriendDao {
	return FriendDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (dao FriendDaoImpl) NewFriend(friend *Friend) (*Friend, error) {
	panic("implement me")
}

func (dao FriendDaoImpl) GetFriendInfo(friendId int64) (*Friend, error) {
	panic("implement me")
}

func (dao FriendDaoImpl) DeleteFriend(cmd DeleteFriendCmd) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (dao FriendDaoImpl) QueryFriendList(userId int64) ([]*Friend, error) {
	//TODO implement me
	panic("implement me")
}

func (dao FriendDaoImpl) QueryFriendInfo(userId int64, friendId int64) (*Friend, error) {
	var sql = fmt.Sprintf("select %s from %s where user_id = ? and friend_id = ?", friendColumns, friendTableName)
	one, err := dao.DatabaseTpl.FindOne(context.Background(), sql, Friend{}, userId, friendId)
	if err != nil || one == nil {
		return nil, err
	}
	friend := one.(Friend)
	return &friend, nil
}
