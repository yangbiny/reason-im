package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/application"
	"reason-im/internal/config/mysql"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
)

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	datasource := mysql.Datasource()

	// user
	dao := &rpcclient.UserDaoImpl{
		DatabaseTpl: &mysql2.DatabaseTpl{
			Db: datasource,
		},
	}
	service := application.NewUserService(dao)
	userApi := NewUserApi(&service)
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", userApi.RegisterNewUser)
		userGroup.GET("/query/", userApi.QueryUserById)
	}

	// friend
	friendDao := &rpcclient.FriendDaoImpl{
		DatabaseTpl: &mysql2.DatabaseTpl{
			Db: datasource,
		},
	}
	inviteService := application.NewFriendInviteService(friendDao)
	friendService := application.NewFriendService(friendDao)
	friendApi := NewFriendApi(friendService)
	friendInviteApi := NewFriendInviteApi(inviteService)
	friendGroup := engine.Group("/friend")
	{
		friendGroup.POST("/invite/", friendInviteApi.InviteFriend)
		friendGroup.POST("/invite/receive/", friendInviteApi.ReceiveInvite)
		friendGroup.POST("/invite/reject/", friendInviteApi.RejectInvite)

		friendGroup.POST("/delete/", friendApi.DeleteFriend)
		friendGroup.POST("/query/list/", friendApi.QueryFriends)
	}

	return engine
}
