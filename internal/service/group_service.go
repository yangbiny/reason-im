package service

import (
	"github.com/gin-gonic/gin"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/repo"
	"reason-im/pkg/model"
)

type Group = model.Group

type GroupService struct {
	groupDao repo.GroupDao
}

func NewGroupService(groupDao repo.GroupDao) GroupService {
	return GroupService{
		groupDao: groupDao,
	}
}

func (service GroupService) NewGroup(ctx *gin.Context, cmd *CreateGroupCmd) (*Group, *apierror.ApiError) {
	group := &Group{
		Name:           cmd.Name,
		Description:    "",
		GroupType:      model.GroupType(cmd.Type),
		GroupMemberCnt: 1,
	}
	newGroup, err := service.groupDao.NewGroup(ctx, group)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	return newGroup, nil

}

type CreateGroupCmd struct {
	UserId int64  `login_user_id:"user_id"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
}
