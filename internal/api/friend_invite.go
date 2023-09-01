package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/application"
	"reason-im/internal/utils/caller"
)

type FriendInviteApi struct {
	FriendService *application.FriendInviteService
}

func NewFriendInviteApi(friendService *application.FriendInviteService) *FriendInviteApi {
	return &FriendInviteApi{
		FriendService: friendService,
	}
}

func (friendInviteApi *FriendInviteApi) QueryFriendsInvite(c *gin.Context) {
	panic("not implemented")
}

func (friendInviteApi *FriendInviteApi) InviteFriend(c *gin.Context) {
	var cmd application.InviteFriendCmd
	caller.CallWithCmd(friendInviteApi.FriendService.InviteFriend, c, cmd)
}

func (friendInviteApi *FriendInviteApi) ReceiveInvite(c *gin.Context) {

}

func (friendInviteApi *FriendInviteApi) RejectInvite(c *gin.Context) {

}
