package repo

import (
	"reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

type Group = model.Group

type GroupDao interface {
	NewGroup(group *Group) (*Group, error)
}

type GroupDaoImpl struct {
	databasesTpl *mysql.DatabaseTpl
}

func NewGroupDao(databases *mysql.DatabaseTpl) GroupDao {
	return GroupDaoImpl{
		databasesTpl: databases,
	}
}

func (g GroupDaoImpl) NewGroup(group *Group) (*Group, error) {
	//TODO implement me
	panic("implement me")
}
