//go:build wireinject

package reason_im

import (
	"github.com/google/wire"
	"reason-im/internal/config/mysql"
	"reason-im/internal/repo"
	service2 "reason-im/internal/service"
	mysql2 "reason-im/internal/utils/mysql"
)

var datasource = mysql.Datasource()
var tpl = mysql2.NewDatabaseTpl(datasource)

func InitUserService() service2.UserService {
	wire.Build(
		service2.NewUserService,
		repo.NewUserDao,
		repo.NewFriendDao,
		wire.Value(tpl),
	)
	return service2.UserService{}
}

func InitInviteFriendService() service2.FriendInviteService {
	wire.Build(
		service2.NewFriendInviteService,
		repo.NewFriendInviteDao,
		repo.NewFriendDao,
		repo.NewUserDao,
		wire.Value(tpl),
	)
	return service2.FriendInviteService{}
}

func InitFriendService() service2.FriendService {
	wire.Build(
		service2.NewFriendService,
		repo.NewFriendDao,
		wire.Value(tpl),
	)
	return service2.FriendService{}
}

func InitGroupService() service2.GroupService {
	wire.Build(
		service2.NewGroupService,
		repo.NewGroupDao,
		repo.NewUserDao,
		repo.NewGroupMemberDao,
		repo.NewGroupMemberInviteDao,
		wire.Value(tpl),
	)
	return service2.GroupService{}
}

func InitMsgService() service2.MsgService {
	wire.Build(
		service2.NewMsgService,
		repo.NewFriendDao,
		repo.NewGroupMemberDao,
		repo.NewGroupDao,
		wire.Value(tpl),
	)
	return service2.MsgService{}
}

func InitSubService() service2.SubService {
	wire.Build(
		service2.NewSubService,
		service2.NewMsgService,
		repo.NewFriendDao,
		repo.NewGroupMemberDao,
		repo.NewGroupDao,
		wire.Value(tpl),
	)
	return service2.SubService{}
}
