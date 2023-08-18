package rpcclient

import (
	"reason-im/internal/config/mysql"
)

type User struct {
	Id   int64
	Name string
}

type UserClient interface {
	NewUser(user *User) *User
	GetUserInfo(userId int64) *User
}

type UserClientHandler struct {
	Client      UserClient
	MysqlClient mysql.MysqlConfig
}

func (u UserClientHandler) NewUser(user *User) *User {

	panic("implement me")
}

func (u UserClientHandler) GetUserInfo(userId int64) *User {
	//TODO implement me
	panic("implement me")
}
