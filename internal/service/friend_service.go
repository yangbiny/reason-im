package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	slice_utils "github.com/yangbiny/reason-commons/slice"
	"reason-im/internal/repo"
	"reason-im/internal/utils/logger"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/dto/vo"
	"reason-im/pkg/model"
	"time"
)

type FriendService struct {
	friendDao repo.FriendDao
}

type FriendInviteService struct {
	friendInviteDao repo.FriendInviteDao
	friendDao       repo.FriendDao
	userDao         repo.UserDao
	databaseTpl     *mysql2.DatabaseTpl
}

type FriendInvite = model.FriendInvite

type Friend = model.Friend

func NewFriendService(friendDao repo.FriendDao) FriendService {
	return FriendService{
		friendDao: friendDao,
	}
}

func NewFriendInviteService(
	friendInviteDao repo.FriendInviteDao,
	friendDao repo.FriendDao,
	userDao repo.UserDao,
	tpl *mysql2.DatabaseTpl,
) FriendInviteService {
	return FriendInviteService{
		friendInviteDao: friendInviteDao,
		friendDao:       friendDao,
		userDao:         userDao,
		databaseTpl:     tpl,
	}
}

func (service *FriendInviteService) InviteFriend(ctx *gin.Context, cmd *InviteFriendCmd) (*FriendInvite, error) {
	if cmd.UserId == cmd.FriendId {
		return nil, errors.WithStack(fmt.Errorf("不能添加自己为好友"))
	}
	userInfo, err := service.userDao.GetUserInfo(cmd.UserId)
	if err != nil {
		logger.ErrorWithErr(context.Background(), "query has error", err)
		return nil, err
	}
	if userInfo == nil {
		logger.Err(context.Background(), "can not find login user_vo Id : %d", cmd.UserId)
		return nil, errors.WithStack(fmt.Errorf("找不到用户信息：%d", cmd.UserId))
	}
	friendUserInfo, _ := service.userDao.GetUserInfo(cmd.FriendId)
	if friendUserInfo == nil {
		logger.Err(context.Background(), "can not find friend user_vo Id :%d ", cmd.FriendId)
		return nil, fmt.Errorf("找不到用户信息：%d", cmd.FriendId)
	}
	friendInfo, err := service.friendDao.QueryFriendInfo(cmd.UserId, cmd.FriendId)
	if err != nil {
		return nil, err
	}
	if friendInfo != nil {
		return nil, fmt.Errorf("该用户已经是你的好友了")
	}

	info, err := service.friendDao.QueryFriendInfo(cmd.UserId, cmd.FriendId)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return nil, fmt.Errorf("该用户已经是你的好友了")
	}
	inviteInfo, err := service.friendInviteDao.GetFriendInviteInfo(context.Background(), cmd.UserId, cmd.FriendId)
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
	inviteInfo, err = service.friendInviteDao.NewFriend(context.Background(), model.NewFriendInvite(cmd.UserId, cmd.FriendId))
	if err != nil {
		return nil, err
	}
	return inviteInfo, nil
}

func (service *FriendInviteService) ReceiveInvite(ctx *gin.Context, cmd *ReceiveInviteCmd) (bool, error) {
	background := context.Background()
	invite, err := service.friendInviteDao.QueryInvite(background, cmd.InviteId)
	if err != nil {
		return false, err
	}
	if invite == nil || invite.FriendId != cmd.UserId {
		return false, errors.WithStack(fmt.Errorf("邀请信息不存在"))
	}
	if invite.Status != model.INVITE {
		return false, errors.WithStack(fmt.Errorf("邀请信息已经处理过了"))
	}

	var friend1 = Friend{
		UserId:    invite.UserId,
		FriendId:  invite.FriendId,
		Remark:    invite.Extra,
		Status:    model.NORMAL,
		GmtCreate: time.Now(),
		GmtUpdate: time.Now(),
	}

	var friend2 = Friend{
		UserId:    invite.FriendId,
		FriendId:  invite.UserId,
		Remark:    "",
		Status:    model.NORMAL,
		GmtCreate: time.Now(),
		GmtUpdate: time.Now(),
	}
	invite.ReceiveInvite()

	err = service.databaseTpl.WithTransaction(&background, func(tx mysql2.Transaction) error {
		_, err = service.friendInviteDao.UpdateInvite(background, invite)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		_, err1 := service.friendDao.NewFriend(&friend1)
		_, err2 := service.friendDao.NewFriend(&friend2)
		if err1 != nil || err2 != nil {
			_ = tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service *FriendInviteService) RejectInvite(ctx *gin.Context, cmd *RejectInviteCmd) (bool, error) {
	background := context.Background()
	invite, err := service.friendInviteDao.QueryInvite(background, cmd.InviteId)
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("邀请信息不存在"))
	}
	if invite == nil || invite.UserId != cmd.UserId {
		return false, nil
		//return false, errors.WithStack(fmt.Errorf("邀请信息不存在"))
	}
	if invite.Status != model.INVITE {
		return false, errors.WithStack(fmt.Errorf("邀请信息已经处理过了"))
	}
	invite.RejectInvite()
	_, err = service.friendInviteDao.UpdateInvite(background, invite)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service *FriendInviteService) QueryInviteList(ctx *gin.Context, cmd *QueryInviteCmd) ([]*vo.UserInviteVo, error) {
	list, err := service.friendInviteDao.QueryBeInviteFriendList(ctx, cmd.UserId)
	if err != nil {
		return nil, err
	}
	slice_utils.MapTo(list, func(invite *FriendInvite) int64 {
		return invite.UserId
	})
	var result []*vo.UserInviteVo
	for _, invite := range list {
		info, err := service.userDao.GetUserInfo(invite.UserId)
		if err != nil {
			logger.ErrorWithErr(ctx, "query has error", err)
			continue
		}
		if info == nil {
			continue
		}
		result = append(result, &vo.UserInviteVo{
			Id:             invite.Id,
			InviteUserId:   invite.UserId,
			InviteUserName: info.Name,
			InviteStatus:   int(invite.Status),
			Extra:          invite.Extra,
			GmtCreate:      invite.GmtCreate,
		})
	}
	return result, nil
}

func (service *FriendService) DeleteFriend(ctx *gin.Context, cmd *DeleteFriendCmd) (bool, error) {
	panic("")
}

func (service *FriendService) QueryFriends(ctx *gin.Context, cmd *QueryFriendCmd) ([]*Friend, error) {
	panic("")
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

type RejectInviteCmd struct {
	InviteId int64 `json:"invite_id"`
	UserId   int64 `login_user_id:"user_id"`
}

type QueryInviteCmd struct {
	UserId int64 `login_user_id:"user_id" required:"true"`
}

type QueryFriendCmd struct {
}

type DeleteFriendCmd struct {
	UserId   int64 `login_user_id:"user_id"`
	FriendId int64 `json:"friend_id" binding:"required"`
}
