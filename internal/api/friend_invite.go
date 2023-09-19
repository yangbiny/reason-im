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
	caller.Call(friendInviteApi.FriendService.InviteFriend, c, new(service.InviteFriendCmd))
}

func (friendInviteApi *FriendInviteApi) ReceiveInvite(c *gin.Context) {
	caller.Call(friendInviteApi.FriendService.ReceiveInvite, c, new(service.ReceiveInviteCmd))
}

func (friendInviteApi *FriendInviteApi) RejectInvite(c *gin.Context) {

}
