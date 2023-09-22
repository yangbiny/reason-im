package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	reason_im "reason-im"
	"reason-im/internal/config/web"
	"reason-im/internal/utils/caller"
	"reason-im/internal/utils/logger"
	"reason-im/pkg/service"
)

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.ErrorWithErr(nil, "set trusted proxies has failed", errors.WithStack(err))
		return nil
	}
	// user
	userService := reason_im.InitUserService()
	userApi := NewUserApi(&userService)
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", userApi.RegisterNewUser)
		userGroup.GET("/query/", userApi.QueryUserById)
		userGroup.POST("/login/", userApi.Login)
	}

	// friend
	inviteService := reason_im.InitInviteFriendService()
	friendService := reason_im.InitFriendService()
	friendApi := NewFriendApi(friendService)
	friendInviteApi := NewFriendInviteApi(inviteService)
	friendGroup := engine.Group("/friend", web.Authorize())
	{
		friendGroup.POST("/invite/", onEvent(new(service.InviteFriendCmd), friendInviteApi.InviteFriendService.InviteFriend))
		friendGroup.POST("/invite/receive/", friendInviteApi.ReceiveInvite)
		friendGroup.POST("/invite/reject/", friendInviteApi.RejectInvite)

		friendGroup.POST("/delete/", friendApi.DeleteFriend)
		friendGroup.POST("/query/list/", friendApi.QueryFriends)
	}

	return engine
}

func onEvent[Req, Resp any](req Req, pairs func(command Req) (Resp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		caller.Call(pairs, c, req)
	}
}
