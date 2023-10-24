package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yangbiny/reason-commons/err"
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

func (service GroupService) NewGroup(ctx *gin.Context, cmd *CreateGroupCmd) (*Group, *err.ApiError) {
	panic("implement me")
}

type CreateGroupCmd struct {
}
