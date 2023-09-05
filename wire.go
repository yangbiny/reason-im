//go:build wireinject

package reason_im

import (
	"github.com/google/wire"
	"reason-im/internal/config/mysql"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
	"reason-im/pkg/service"
)

var datasource = mysql.Datasource()
var tpl = mysql2.NewDatabaseTpl(datasource)

func InitUserService() service.UserService {
	wire.Build(
		service.NewUserService,
		rpcclient.NewUserDao,
		wire.Value(tpl),
	)
	return service.UserService{}
}

func InitInviteFriendService() service.FriendInviteService {
	wire.Build(
		service.NewFriendInviteService,
		rpcclient.NewFriendInviteDao,
		wire.Value(tpl),
	)
	return service.FriendInviteService{}
}
