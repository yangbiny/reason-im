package api

import (
	"github.com/gin-gonic/gin"
	reason_im "reason-im"
	"reason-im/internal/config/mysql"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/rpcclient"
)

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	datasource := mysql.Datasource()
	// user
	userService := reason_im.InitUserService()
	userApi := NewUserApi(&userService)
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
	inviteService := userService.NewFriendInviteService(friendDao)
	friendService := userService.NewFriendService(friendDao)
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
