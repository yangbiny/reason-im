package repo

import (
	"context"
	"fmt"
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
	"time"
)

type Group = model.Group
type GroupMember = model.GroupMember

var columns = "id,name,description,group_type,group_member_cnt,gmt_create,gmt_update"
var groupMemberColumns = "id,group_id,user_id,nick_name,group_member_role,gmt_create,gmt_update"
var imGroupTableName = "im_group"
var imGroupMemberTableName = "im_group_member"

type GroupDao interface {
	NewGroup(ctx *context.Context, group *Group) (*Group, error)
	FindById(ctx *context.Context, id int64) (*Group, error)
}

type GroupMemberDao interface {
	NewGroupMember(ctx *context.Context, groupMember *GroupMember) (*GroupMember, error)
	FindByGroupId(ctx *context.Context, groupId int64) ([]*GroupMember, error)
}

type GroupDaoImpl struct {
	databasesTpl *mysql.DatabaseTpl
}

type GroupMemberDaoImpl struct {
	databasesTpl *mysql.DatabaseTpl
}

func NewGroupMemberDao(databases *mysql.DatabaseTpl) GroupMemberDao {
	return &GroupMemberDaoImpl{
		databasesTpl: databases,
	}
}

func NewGroupDao(databases *mysql.DatabaseTpl) GroupDao {
	return &GroupDaoImpl{
		databasesTpl: databases,
	}
}

func (g *GroupMemberDaoImpl) FindByGroupId(ctx *context.Context, groupId int64) ([]*GroupMember, error) {
	sqlStr := fmt.Sprintf("select %s from %s where group_id = ?", groupMemberColumns, imGroupMemberTableName)
	list, err := g.databasesTpl.FindList(ctx, sqlStr, GroupMemberDO{}, groupId)
	if err != nil {
		return nil, err
	}
	var groupMembers []*GroupMember
	for _, item := range list {
		groupMemberDO := item.(*GroupMemberDO)
		groupMembers = append(groupMembers, &GroupMember{
			Id:              groupMemberDO.Id,
			GroupId:         groupMemberDO.GroupId,
			UserId:          groupMemberDO.UserId,
			NickName:        groupMemberDO.NickName,
			GroupMemberRole: model.GroupMemberRole(groupMemberDO.GroupMemberRole),
			GmtCreate:       groupMemberDO.GmtCreate,
			GmtUpdate:       groupMemberDO.GmtUpdate,
		})
	}
	return groupMembers, nil
}

func (g *GroupMemberDaoImpl) NewGroupMember(ctx *context.Context, groupMember *GroupMember) (*GroupMember, error) {
	sqlStr := fmt.Sprintf("insert into %s (group_id,user_id,nick_name,group_member_role,gmt_create,gmt_update) values (?,?,?,?,?,?)", imGroupMemberTableName)
	insert, err := g.databasesTpl.Insert(ctx, sqlStr, groupMember.GroupId, groupMember.UserId, groupMember.NickName, groupMember.GroupMemberRole, groupMember.GmtCreate, groupMember.GmtUpdate)
	if err != nil {
		return nil, err
	}
	groupMember.Id = insert
	return groupMember, nil
}

func (g *GroupDaoImpl) NewGroup(ctx *context.Context, group *Group) (*Group, error) {
	sqlStr := fmt.Sprintf("insert into %s (name,description,group_type,group_member_cnt,gmt_create,gmt_update) values (?,?,?,?,?,?)", imGroupTableName)
	insert, err := g.databasesTpl.Insert(ctx, sqlStr, group.Name, group.Description, group.GroupType, group.GroupMemberCnt, group.GmtCreate, group.GmtUpdate)
	if err != nil {
		return nil, err
	}
	group.Id = insert
	return group, nil
}

func (g *GroupDaoImpl) FindById(ctx *context.Context, id int64) (*Group, error) {
	sqlStr := fmt.Sprintf("select %s from %s where id = ?", columns, imGroupTableName)
	one, err := g.databasesTpl.FindOne(ctx, sqlStr, GroupDO{}, id)
	if err != nil {
		return nil, err
	}
	groupDo := one.(*GroupDO)
	return &Group{
		Id:             groupDo.Id,
		Name:           groupDo.Name,
		Description:    groupDo.Description,
		GroupType:      model.GroupType(groupDo.GroupType),
		GroupMemberCnt: groupDo.GroupMemberCnt,
		GmtCreate:      groupDo.GmtCreate,
		GmtUpdate:      groupDo.GmtUpdate,
	}, nil
}

type GroupDO struct {
	Id             int64     `mysql:"id"`
	Name           string    `mysql:"name"`
	Description    string    `mysql:"description"`
	GroupType      int       `mysql:"group_type"`
	GroupMemberCnt int32     `mysql:"group_member_cnt"`
	GmtCreate      time.Time `mysql:"gmt_create"`
	GmtUpdate      time.Time `mysql:"gmt_update"`
}

type GroupMemberDO struct {
	Id              int64     `mysql:"id"`
	GroupId         int64     `mysql:"group_id"`
	UserId          int64     `mysql:"user_id"`
	NickName        string    `mysql:"nick_name"`
	GroupMemberRole int       `mysql:"group_member_role"`
	GmtCreate       time.Time `mysql:"gmt_create"`
	GmtUpdate       time.Time `mysql:"gmt_update"`
}
