package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/service"
)

type FriendInviteApi struct {
	FriendService service.FriendInviteService
}

func NewFriendInviteApi(friendService service.FriendInviteService) *FriendInviteApi {
	return &FriendInviteApi{
		FriendService: friendService,
	}
}

func (friendInviteApi *FriendInviteApi) QueryFriendsInvite(c *gin.Context) {
	panic("not implemented")
}

func (friendInviteApi *FriendInviteApi) InviteFriend(c *gin.Context) {
	var cmd service.InviteFriendCmd
	caller.CallWithCmd(friendInviteApi.FriendService.InviteFriend, c, cmd)
}

func (friendInviteApi *FriendInviteApi) ReceiveInvite(c *gin.Context) {

}

func (friendInviteApi *FriendInviteApi) RejectInvite(c *gin.Context) {

}
