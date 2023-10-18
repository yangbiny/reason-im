// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package reason_im

import (
	"reason-im/internal/config/mysql"
	"reason-im/internal/repo"
	"reason-im/internal/service"
	mysql2 "reason-im/internal/utils/mysql"
)

// Injectors from wire.go:

func InitUserService() service.UserService {
	databaseTpl := _wireDatabaseTplValue
	userDao := repo.NewUserDao(databaseTpl)
	userService := service.NewUserService(userDao)
	return userService
}

var (
	_wireDatabaseTplValue = tpl
)

func InitInviteFriendService() service.FriendInviteService {
	databaseTpl := _wireMysqlDatabaseTplValue
	friendInviteDao := repo.NewFriendInviteDao(databaseTpl)
	friendDao := repo.NewFriendDao(databaseTpl)
	userDao := repo.NewUserDao(databaseTpl)
	friendInviteService := service.NewFriendInviteService(friendInviteDao, friendDao, userDao, databaseTpl)
	return friendInviteService
}

var (
	_wireMysqlDatabaseTplValue = tpl
)

func InitFriendService() service.FriendService {
	databaseTpl := _wireDatabaseTplValue2
	friendDao := repo.NewFriendDao(databaseTpl)
	friendService := service.NewFriendService(friendDao)
	return friendService
}

var (
	_wireDatabaseTplValue2 = tpl
)

// wire.go:

var datasource = mysql.Datasource()

var tpl = mysql2.NewDatabaseTpl(datasource)
