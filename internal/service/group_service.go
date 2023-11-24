package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/repo"
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/dto/vo"
	"reason-im/pkg/model"
	"time"
)

type Group = model.Group

type GroupService struct {
	groupDao             repo.GroupDao
	groupMemberDao       repo.GroupMemberDao
	groupMemberInviteDao repo.GroupMemberInviteDao
	userDao              repo.UserDao
	tpl                  *mysql.DatabaseTpl
}

func NewGroupService(
	groupDao repo.GroupDao,
	userDao repo.UserDao,
	groupMemberDao repo.GroupMemberDao,
	dao repo.GroupMemberInviteDao,
	tpl *mysql.DatabaseTpl,
) GroupService {
	return GroupService{
		groupDao:             groupDao,
		groupMemberDao:       groupMemberDao,
		groupMemberInviteDao: dao,
		userDao:              userDao,
		tpl:                  tpl,
	}
}

func (service GroupService) NewGroup(ctx *gin.Context, cmd *CreateGroupCmd) (*vo.GroupVo, *apierror.ApiError) {
	now := time.Now()
	group := &Group{
		Name:           cmd.Name,
		Description:    "",
		GroupType:      model.GroupType(cmd.Type),
		GroupMemberCnt: 1,
		GmtCreate:      now,
		GmtUpdate:      now,
	}
	var newGroup *Group
	var err error
	ctx2 := ctx.Request.Context()
	err = service.tpl.WithTransaction(ctx2, func(ctx context.Context) error {
		newGroup, err = service.groupDao.NewGroup(ctx, group)
		if err != nil {
			return err
		}
		info, err2 := service.userDao.GetUserInfo(ctx, cmd.UserId)
		if err2 != nil {
			return err2
		}
		groupMember := &model.GroupMember{
			GroupId:         newGroup.Id,
			UserId:          cmd.UserId,
			NickName:        info.Name,
			GroupMemberRole: model.GROUP_MEMBER_ROLE_OWNER,
			GmtCreate:       now,
			GmtUpdate:       now,
		}
		_, err = service.groupMemberDao.NewGroupMember(ctx, groupMember)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}

	return &vo.GroupVo{
		Id:   newGroup.Id,
		Name: newGroup.Name,
	}, nil
}

func (service GroupService) InviteToGroup(ctx *gin.Context, cmd *InviteUserToGroupCmd) (*vo.InviteGroupVo, *apierror.ApiError) {
	ctx2 := ctx.Request.Context()
	id, err := service.groupDao.FindById(ctx2, cmd.GroupId)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	if id == nil {
		return nil, apierror.WhenParamError(fmt.Errorf("群组不存在"))
	}
	groupMember, err := service.groupMemberDao.FindByGroupAndUserId(ctx2, cmd.GroupId, cmd.UserId)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	if !groupMember.CanInviteUser() {
		return nil, apierror.WhenParamError(fmt.Errorf("你没有权限邀请用户"))
	}
	inviteUser, _ := service.userDao.GetUserInfo(ctx2, cmd.InviteUserId)
	if inviteUser == nil {
		return nil, apierror.WhenParamError(fmt.Errorf("邀请用户不存在"))
	}
	member, _ := service.groupMemberDao.FindByGroupAndUserId(ctx, cmd.GroupId, cmd.InviteUserId)
	if member != nil {
		return nil, apierror.WhenParamError(fmt.Errorf("该用户已经在该群组中"))
	}
	invite, _ := service.groupMemberInviteDao.NewGroupMemberInvite(ctx, &model.GroupMemberInvite{})
	return &vo.InviteGroupVo{
		Id:           invite.Id,
		GroupId:      invite.GroupId,
		InviteUserId: invite.InviteUserId,
	}, nil
}

func (service GroupService) SendMsgToGroup(ctx *gin.Context, cmd *GroupMemberSendMsgCmd) (bool, *apierror.ApiError) {
	ctx2 := ctx.Request.Context()
	group, err := service.groupDao.FindById(ctx2, cmd.GroupId)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}
	if group == nil {
		return false, apierror.WhenParamError(fmt.Errorf("群组不存在"))
	}
	id, err := service.groupMemberDao.FindByGroupAndUserId(ctx2, cmd.GroupId, cmd.UserId)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}
	if id == nil {
		return false, apierror.WhenParamError(fmt.Errorf("你不在该群组中"))
	}

	groupMembers, err := service.groupMemberDao.FindByGroupId(ctx2, cmd.GroupId)
	if len(groupMembers) <= 0 {
		return false, apierror.WhenParamError(fmt.Errorf("群组中没有用户"))
	}

	for _, member := range groupMembers {
		msg := Msg{
			MsgType:    model.MsgTypeFriend,
			Msg:        &cmd.Msg,
			FromUserId: &cmd.UserId,
			ToId:       &cmd.GroupId,
			Ext:        nil,
		}
		SendMsg(&member.UserId, &msg)
	}
	return false, nil
}

type GroupMemberSendMsgCmd struct {
	UserId  int64  `login_user_id:"user_id"`
	GroupId int64  `json:"group_id"`
	Msg     string `json:"msg"`
}

type CreateGroupCmd struct {
	UserId int64  `login_user_id:"user_id"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
}

type InviteUserToGroupCmd struct {
	UserId       int64 `login_user_id:"user_id"`
	GroupId      int64 `json:"group_id"`
	InviteUserId int64 `json:"invite_user_id"`
}
