package rpcclient

import (
	"context"
	"fmt"
	"reason-im/internal/config/mysql"
	"reason-im/internal/utils/logger"
)

var tableName = "im_user"

type User struct {
	Id   int64
	Name string
}

type UserClient interface {
	NewUser(user *User) *User
	GetUserInfo(userId int64) *User
}

type UserClientHandler struct {
	Client UserClient
}

func (u UserClientHandler) NewUser(user *User) *User {
	ctx := context.Background()
	connection := mysql.GetConnection(ctx)
	defer mysql.CloseMysqlConn(connection, ctx)
	var sqlStr = fmt.Sprintf("insert into %s (name) values (?)", tableName)
	prepareContext, _ := connection.PrepareContext(ctx, sqlStr)
	result, err := prepareContext.Exec(user.Name)
	if err != nil {
		logger.Error(ctx, "execute sql has failed", "sql", sqlStr, "username", user.Name)
		panic("execute has failed")
	}
	id, _ := result.LastInsertId()
	return &User{
		Id:   id,
		Name: user.Name,
	}
}

func (u UserClientHandler) GetUserInfo(userId int64) *User {
	//TODO implement me
	panic("implement me")
}
