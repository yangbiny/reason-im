package rpcclient

import (
	"context"
	"fmt"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
	"time"
)

type User model.Users

var tableName = "im_user"

type UserDao interface {
	NewUser(user *User) *User
	GetUserInfo(userId int64) *User
}

type UserDaoImpl struct {
	DatabaseTpl *mysql2.DatabaseTpl
}

func NewUserDao(tpl *mysql2.DatabaseTpl) UserDao {
	return UserDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (u UserDaoImpl) NewUser(user *User) *User {
	ctx := context.Background()
	var sqlStr = fmt.Sprintf("insert into %s (name,gmt_create,gmt_update) values (?,?,?)", tableName)
	date := time.Now().Local()
	insert := u.DatabaseTpl.Insert(ctx, sqlStr, user.Name, date, date)
	return &User{
		Id:   insert,
		Name: user.Name,
	}
}

func (u UserDaoImpl) GetUserInfo(userId int64) *User {
	ctx := context.Background()
	sql := fmt.Sprintf("select * from %s where id = ?", tableName)
	user := u.DatabaseTpl.FindOne(ctx, sql, User{}, userId).(User)
	return &user
}
