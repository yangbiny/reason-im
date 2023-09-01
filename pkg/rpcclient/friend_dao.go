package rpcclient

import (
	"reason-im/internal/entity"
	mysql2 "reason-im/internal/utils/mysql"
)

type Friend entity.Friend

type DeleteFriendCmd struct {
	Id     int64
	Status entity.FriendStatus
}

type FriendClient interface {
	NewFriend(friend *Friend) *Friend
	GetFriendInfo(friendId int64) *Friend
	DeleteFriend(cmd DeleteFriendCmd) bool
	QueryFriendList(userId int64) []*Friend
	QueryFriendInfo(userId int64, friendId int64) *Friend
}

type FriendDaoImpl struct {
	DatabaseTpl *mysql2.DatabaseTpl
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
	//TODO implement me
	panic("implement me")
}
