// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package reason_im

import (
	"reason-im/internal/config/mysql"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
	"reason-im/pkg/service"
)

// Injectors from wire.go:

func InitUserService() service.UserService {
	databaseTpl := _wireDatabaseTplValue
	userDao := rpcclient.NewUserDao(databaseTpl)
	userService := service.NewUserService(userDao)
	return userService
}

var (
	_wireDatabaseTplValue = tpl
)

func InitInviteFriendService() service.FriendInviteService {
	databaseTpl := _wireMysqlDatabaseTplValue
	friendInviteDao := rpcclient.NewFriendInviteDao(databaseTpl)
	friendInviteService := service.NewFriendInviteService(friendInviteDao)
	return friendInviteService
}

var (
	_wireMysqlDatabaseTplValue = tpl
)

// wire.go:

var datasource = mysql.Datasource()

var tpl = mysql2.NewDatabaseTpl(datasource)
