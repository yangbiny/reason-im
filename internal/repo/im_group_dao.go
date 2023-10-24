package repo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

type Group = model.Group

var columns = "id,name,description,group_type,group_member_cnt,gmt_create,gmt_update"
var imGroupTableName = "im_group"

type GroupDao interface {
	NewGroup(ctx *gin.Context, group *Group) (*Group, error)
	FindById(ctx *gin.Context, id int64) (*Group, error)
}

type GroupDaoImpl struct {
	databasesTpl *mysql.DatabaseTpl
}

func NewGroupDao(databases *mysql.DatabaseTpl) GroupDao {
	return GroupDaoImpl{
		databasesTpl: databases,
	}
}

func (g GroupDaoImpl) FindById(ctx *gin.Context, id int64) (*Group, error) {
	sqlStr := fmt.Sprintf("select %s from %s where id = ?", columns, imGroupTableName)
	one, err := g.databasesTpl.FindOne(ctx, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return one.(*Group), nil
}

func (g GroupDaoImpl) NewGroup(ctx *gin.Context, group *Group) (*Group, error) {
	sqlStr := fmt.Sprintf("insert into %s (name,description,group_type,group_member_cnt,gmt_create,gmt_update) values (?,?,?,?,?,?)", imGroupTableName)
	insert, err := g.databasesTpl.Insert(ctx, sqlStr, group.Name, group.Description, group.GroupType, group.GroupMemberCnt, group.GmtCreate, group.GmtUpdate)
	if err != nil {
		return nil, err
	}
	group.Id = insert
	return group, nil
}
