package service

import (
	"reason-im/pkg/model"
	"reason-im/pkg/rpcclient"
)

type FriendService struct {
	friendDao rpcclient.FriendDao
}

type FriendInviteService struct {
	friendInviteDao rpcclient.FriendInviteDao
	friendDao       *rpcclient.FriendDaoImpl
	userDao         *rpcclient.UserDaoImpl
}

type FriendInvite model.FriendInvite

func NewFriendService(friendDao rpcclient.FriendDao) FriendService {
	return FriendService{
		friendDao: friendDao,
	}
}

func NewFriendInviteService(friendInviteDao rpcclient.FriendInviteDao) FriendInviteService {
	return FriendInviteService{
		friendInviteDao: friendInviteDao,
	}
}

func (service *FriendService) QueryUserFriend() {

}

func (service *FriendInviteService) InviteFriend(cmd InviteFriendCmd) FriendInvite {
	info := service.friendDao.QueryFriendInfo(cmd.UserId, cmd.FriendId)
	if info != nil {
		panic("该用户已经是你的好友了")
	}
	panic("")
}

type InviteFriendCmd struct {
	UserId   int64  `login_user_id:"user_id"`
	FriendId int64  `json:"friend_id"`
	Remark   string `json:"remark"`
}
