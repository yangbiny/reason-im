package service

import (
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

func NewGroupService(groupDao repo.GroupDao, tpl *mysql.DatabaseTpl) GroupService {
	return GroupService{
		groupDao: groupDao,
		tpl:      tpl,
	}
}

func (service GroupService) NewGroup(ctx *gin.Context, cmd *CreateGroupCmd) (*vo.GroupVo, *apierror.ApiError) {
	group := &Group{
		Name:           cmd.Name,
		Description:    "",
		GroupType:      model.GroupType(cmd.Type),
		GroupMemberCnt: 1,
	}
	var newGroup *Group
	var err error
	err = service.tpl.WithTransaction(ctx, func(tx mysql.Transaction) error {
		defer func() {
			if err := recover(); err != nil {
				_ = tx.Rollback()
			}
		}()
		newGroup, err = service.groupDao.NewGroup(ctx, group)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		info, err2 := service.userDao.GetUserInfo(cmd.UserId)
		if err2 != nil {
			_ = tx.Rollback()
			return err2
		}
		groupMember := &model.GroupMember{
			GroupId:         newGroup.Id,
			UserId:          cmd.UserId,
			NickName:        info.Name,
			GroupMemberRole: model.GROUP_MEMBER_ROLE_ADMIN,
			GmtCreate:       time.Time{},
			GmtUpdate:       time.Time{},
		}
		_, err = service.groupMemberDao.NewGroupMember(ctx, groupMember)
		if err != nil {
			_ = tx.Rollback()
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
