package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/yangbiny/reason-commons/err"
	reasonim "reason-im"
	"reason-im/internal/config/web"
	service2 "reason-im/internal/service"
	"reason-im/internal/utils/caller"
	"reason-im/internal/utils/logger"
)

func NewGinRouter() *gin.Engine {
	engine := gin.New()
	err2 := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err2 != nil {
		logger.ErrorWithErr(nil, "set trusted proxies has failed", errors.WithStack(err2))
		return nil
	}
	// user
	userService := reasonim.InitUserService()
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register/", onEvent(new(service2.NewUserCmd), userService.NewUser))
		userGroup.POST("/login/", onEvent(new(service2.UserLoginCmd), userService.Login))
	}

	loginUserGroup := engine.Group("/user", web.Authorize())
	{
		loginUserGroup.GET("/query/relation/:query_uid/", onEvent(new(service2.QueryUserCmd), userService.GetUserInfo))
	}

	// friend
	inviteFriendService := reasonim.InitInviteFriendService()
	friendService := reasonim.InitFriendService()
	friendGroup := engine.Group("/friend", web.Authorize())
	{
		friendGroup.POST("/invite/", onEvent(new(service2.InviteFriendCmd), inviteFriendService.InviteFriend))
		friendGroup.POST("/invite/receive/", onEvent(new(service2.ReceiveInviteCmd), inviteFriendService.ReceiveInvite))
		friendGroup.POST("/invite/reject/", onEvent(new(service2.RejectInviteCmd), inviteFriendService.RejectInvite))
		friendGroup.POST("/delete/", onEvent(new(service2.DeleteFriendCmd), friendService.DeleteFriend))
		friendGroup.POST("/query/list/", onEvent(new(service2.QueryFriendCmd), friendService.QueryFriends))

		friendGroup.GET("/invite/list/", onEvent(new(service2.QueryInviteCmd), inviteFriendService.QueryInviteList))
	}

	groupService := reasonim.InitGroupService()
	imGroup := engine.Group("/group", web.Authorize())
	{
		imGroup.POST("/create/", onEvent(new(service2.CreateGroupCmd), groupService.NewGroup))
		imGroup.POST("/invite/", onEvent(new(service2.InviteUserToGroupCmd), groupService.InviteToGroup))
		imGroup.POST("/msg/send/", onEvent(new(service2.GroupMemberSendMsgCmd), groupService.SendMsgToGroup))
	}

	service := reasonim.InitMsgService()
	imMsgGroup := engine.Group("/im/msg/", web.Authorize())
	{
		imMsgGroup.POST("/send/", onEvent(new(service2.MsgCmd), service.SendMsg))
	}

	ws := engine.Group("/ws/", web.Authorize())
	{
		ws.GET("connect/", onEvent(new(service2.MessageCmd), service2.ServeWs))
	}
	return engine
}

func onEvent[Req, Resp any](req Req, pairs func(ctx *gin.Context, command Req) (Resp, *err.ApiError)) gin.HandlerFunc {
	return func(c *gin.Context) {
		caller.Call(pairs, c, req)
	}
}
