package repo

import (
	"context"
	"fmt"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
	"time"
)

type User = model.Users

var tableName = "im_user"

type UserDao interface {
	NewUser(ctx *context.Context, user *User) (*User, error)
	GetUserInfo(ctx *context.Context, userId int64) (*User, error)
	QueryUserByName(ctx *context.Context, name string) (*User, error)
}

type UserDaoImpl struct {
	DatabaseTpl *mysql2.DatabaseTpl
}

func NewUserDao(tpl *mysql2.DatabaseTpl) UserDao {
	return UserDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (u UserDaoImpl) NewUser(ctx *context.Context, user *User) (*User, error) {
	var sqlStr = fmt.Sprintf("insert into %s (name,gmt_create,gmt_update) values (?,?,?)", tableName)
	date := time.Now().Local()
	insert, err := u.DatabaseTpl.Insert(ctx, sqlStr, user.Name, date, date)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:   insert,
		Name: user.Name,
	}, nil
}

func (u UserDaoImpl) GetUserInfo(ctx *context.Context, userId int64) (*User, error) {
	sql := fmt.Sprintf("select * from %s where id = ?", tableName)
	one, err := u.DatabaseTpl.FindOne(ctx, sql, User{}, userId)
	if err != nil || one == nil {
		return nil, err
	}
	user := one.(User)
	return &user, nil
}

func (u UserDaoImpl) QueryUserByName(ctx *context.Context, name string) (*User, error) {
	var sql = "select * from " + tableName + " where name = ?"
	one, err := u.DatabaseTpl.FindOne(ctx, sql, User{}, name)
	if err != nil || one == nil {
		return nil, err
	}
	user := one.(User)
	return &user, nil
}
