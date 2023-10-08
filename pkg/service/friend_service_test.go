package service

import (
	"golang.org/x/net/context"
	"reason-im/internal/config/mysql"
	"reason-im/internal/utils/logger"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
	"testing"
)

func TestFriendInviteService_ReceiveInvite(t *testing.T) {
	datasource := mysql.Datasource()
	tpl := mysql2.NewDatabaseTpl(datasource)
	friendInviteDao := rpcclient.FriendInviteDaoImpl{
		DatabaseTpl: tpl,
	}
	friendDao := rpcclient.FriendDaoImpl{
		DatabaseTpl: tpl,
	}
	userDao := rpcclient.UserDaoImpl{
		DatabaseTpl: tpl,
	}
	var service = NewFriendInviteService(friendInviteDao, friendDao, userDao, tpl)
	cmd := ReceiveInviteCmd{
		UserId:   1,
		InviteId: 2,
	}
	invite, err := service.ReceiveInvite(&cmd)
	if err != nil {
		logger.ErrorWithErr(context.Background(), "query has error", err)
	}
	println(invite)
}
