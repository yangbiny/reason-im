package application

import (
	"reason-im/internal/entity"
	"reason-im/pkg/rpcclient"
)

type FriendService struct {
	friendDao *rpcclient.FriendDaoImpl
}

type FriendInviteService struct {
	friendDao *rpcclient.FriendDaoImpl
	userDao   *rpcclient.UserDaoImpl
}

type FriendInvite entity.FriendInvite

func NewFriendService(friendDao *rpcclient.FriendDaoImpl) *FriendService {
	return &FriendService{
		friendDao: friendDao,
	}
}

func NewFriendInviteService(friendDao *rpcclient.FriendDaoImpl) *FriendInviteService {
	return &FriendInviteService{
		friendDao: friendDao,
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
