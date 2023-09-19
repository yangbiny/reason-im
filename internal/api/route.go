package api

import (
	"github.com/gin-gonic/gin"
	reason_im "reason-im"
	"reason-im/internal/config/web"
)

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	// user
	userService := reason_im.InitUserService()
	userApi := NewUserApi(&userService)
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", userApi.RegisterNewUser)
		userGroup.GET("/query/", userApi.QueryUserById)
	}

	// friend
	inviteService := reason_im.InitInviteFriendService()
	friendService := reason_im.InitFriendService()
	friendApi := NewFriendApi(friendService)
	friendInviteApi := NewFriendInviteApi(inviteService)
	friendGroup := engine.Group("/friend", web.Authorize())
	{
		friendGroup.POST("/invite/", friendInviteApi.InviteFriend)
		friendGroup.POST("/invite/receive/", friendInviteApi.ReceiveInvite)
		friendGroup.POST("/invite/reject/", friendInviteApi.RejectInvite)

		friendGroup.POST("/delete/", friendApi.DeleteFriend)
		friendGroup.POST("/query/list/", friendApi.QueryFriends)
	}

	return engine
}
