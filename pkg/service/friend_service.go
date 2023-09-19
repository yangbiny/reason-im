package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"reason-im/internal/utils/logger"
	"reason-im/pkg/model"
	"reason-im/pkg/rpcclient"
)

type FriendService struct {
	friendDao rpcclient.FriendDao
}

type FriendInviteService struct {
	friendInviteDao rpcclient.FriendInviteDao
	friendDao       rpcclient.FriendDao
	userDao         rpcclient.UserDao
}

type FriendInvite model.FriendInvite

func NewFriendService(friendDao rpcclient.FriendDao) FriendService {
	return FriendService{
		friendDao: friendDao,
	}
}

func NewFriendInviteService(
	friendInviteDao rpcclient.FriendInviteDao,
	friendDao rpcclient.FriendDao,
	userDao rpcclient.UserDao,
) FriendInviteService {
	return FriendInviteService{
		friendInviteDao: friendInviteDao,
		friendDao:       friendDao,
		userDao:         userDao,
	}
}

func (service *FriendService) QueryUserFriend() {

}

func (service *FriendInviteService) InviteFriend(cmd *InviteFriendCmd) (*FriendInvite, error) {
	userInfo, err := service.userDao.GetUserInfo(cmd.UserId)
	if err != nil {
		logger.ErrorWithErr(context.Background(), "query has error", err)
		return nil, err
	}
	if userInfo == nil {
		logger.Err(context.Background(), "can not find login user Id : %d", cmd.UserId)
		return nil, errors.WithStack(fmt.Errorf("找不到用户信息：%d", cmd.UserId))
	}
	friendUserInfo, _ := service.userDao.GetUserInfo(cmd.FriendId)
	if friendUserInfo == nil {
		logger.Err(context.Background(), "can not find friend user Id :%d ", cmd.FriendId)
		return nil, fmt.Errorf("找不到用户信息：%d", cmd.FriendId)
	}
	info, err := service.friendDao.QueryFriendInfo(cmd.UserId, cmd.FriendId)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return nil, fmt.Errorf("该用户已经是你的好友了")
	}
	inviteInfo, err := service.friendInviteDao.GetFriendInviteInfo(cmd.UserId, cmd.FriendId)
	if err != nil {
		return nil, err
	}
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
	inviteInfo, err = service.friendInviteDao.NewFriend(model.NewFriendInvite(cmd.UserId, cmd.FriendId))
	if err != nil {
		return nil, err
	}
	return (*FriendInvite)(inviteInfo), nil
}

func (service *FriendInviteService) ReceiveInvite(cmd *ReceiveInviteCmd) (bool, error) {
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
