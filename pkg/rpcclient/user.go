package rpcclient

import (
	"context"
	"fmt"
	"reason-im/internal/config/mysql"
	"reason-im/internal/utils/logger"
	"time"
)

var tableName = "im_user"

type User struct {
	Id        int64     `mysql:"id"`
	Name      string    `mysql:"name"`
	GmtCreate time.Time `mysql:"gmt_create"`
	GmtUpdate time.Time `mysql:"gmt_update"`
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
	var sqlStr = fmt.Sprintf("insert into %s (name,gmt_create,gmt_update) values (?,?,?)", tableName)
	prepareContext, err := connection.PrepareContext(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
	date := time.Now().Local()
	result, err := prepareContext.Exec(user.Name, date, date)
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
	ctx := context.Background()
	connection := mysql.GetConnection(ctx)
	sql := fmt.Sprintf("select * from %s where id = ?", tableName)
	queryContext, err := connection.QueryContext(ctx, sql, userId)
	if err != nil {
		panic(err)
	}
	if !queryContext.Next() {
		return nil
	}
	var user User
	renderResult := mysql.RenderResult(queryContext, user).(User)
	return &renderResult
}
