package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	reason_im "reason-im"
	"reason-im/internal/config/web"
	service2 "reason-im/internal/service"
	"reason-im/internal/utils/caller"
	"reason-im/internal/utils/logger"
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
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", onEvent(new(service2.NewUserCmd), userService.NewUser))
		userGroup.GET("/query/", onEvent(new(int64), userService.GetUserInfo))
		userGroup.POST("/login/", onEvent(new(service2.UserLoginCmd), userService.Login))
	}

	// friend
	inviteFriendService := reason_im.InitInviteFriendService()
	friendService := reason_im.InitFriendService()
	friendGroup := engine.Group("/friend", web.Authorize())
	{
		friendGroup.POST("/invite/", onEvent(new(service2.InviteFriendCmd), inviteFriendService.InviteFriend))
		friendGroup.POST("/invite/receive/", onEvent(new(service2.ReceiveInviteCmd), inviteFriendService.ReceiveInvite))
		friendGroup.POST("/invite/reject/", onEvent(new(service2.RejectInviteCmd), inviteFriendService.RejectInvite))

		//friendGroup.POST("/delete/", friendApi.DeleteFriend)
		friendGroup.POST("/query/list/", onEvent(new(service2.QueryFriendCmd), friendService.QueryFriends))
	}

	//engine.GET("ws/msg/")
	return engine
}

func onEvent[Req, Resp any](req Req, pairs func(ctx *gin.Context, command Req) (Resp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		caller.Call(pairs, c, req)
	}
}

func onWSRequest(c *gin.Context) {
	caller.CallMS(c, service2.MSService, new(service2.MSServiceCmd))
}
