package service

import (
	"fmt"
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

func (service *FriendInviteService) InviteFriend(cmd InviteFriendCmd) (*FriendInvite, error) {
	info := service.friendDao.QueryFriendInfo(cmd.UserId, cmd.FriendId)
	if info != nil {
		return nil, fmt.Errorf("该用户已经是你的好友了")
	}
	inviteInfo := service.friendInviteDao.GetFriendInviteInfo(cmd.UserId, cmd.FriendId)
	if inviteInfo != nil {
		return &FriendInvite{
			Id:        inviteInfo.Id,
			UserId:    inviteInfo.UserId,
			FriendId:  inviteInfo.FriendId,
			Status:    inviteInfo.Status,
			Extra:     inviteInfo.Extra,
			GmtCreate: inviteInfo.GmtCreate,
			GmtUpdate: inviteInfo.GmtUpdate,
		}, nil
	}
	inviteInfo = service.friendInviteDao.NewFriend(model.NewFriendInvite(cmd.UserId, cmd.FriendId))
	return (*FriendInvite)(inviteInfo), nil
}

func (service *FriendInviteService) ReceiveInvite(cmd ReceiveInviteCmd) (bool, error) {
	panic("implement me")
}

type InviteFriendCmd struct {
	UserId   int64  `login_user_id:"user_id"`
	FriendId int64  `json:"friend_id"`
	Remark   string `json:"remark"`
}

type ReceiveInviteCmd struct {
	InviteId int64 `json:"invite_id"`
	UserId   int64 `login_user_id:"user_id"`
}
