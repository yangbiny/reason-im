package service

import (
	"context"
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
	groupDao       repo.GroupDao
	groupMemberDao repo.GroupMemberDao
	userDao        repo.UserDao
	tpl            *mysql.DatabaseTpl
}

func NewGroupService(
	groupDao repo.GroupDao,
	userDao repo.UserDao,
	groupMemberDao repo.GroupMemberDao,
	tpl *mysql.DatabaseTpl,
) GroupService {
	return GroupService{
		groupDao:       groupDao,
		groupMemberDao: groupMemberDao,
		userDao:        userDao,
		tpl:            tpl,
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
	err = service.tpl.WithTransaction(&ctx2, func(ctx *context.Context) error {
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
	return nil, nil
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
