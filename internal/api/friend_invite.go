package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/caller"
	"reason-im/pkg/service"
)

type FriendInviteApi struct {
	InviteFriendService service.FriendInviteService
}

func NewFriendInviteApi(friendService service.FriendInviteService) *FriendInviteApi {
	return &FriendInviteApi{
		InviteFriendService: friendService,
	}
}

func (friendInviteApi *FriendInviteApi) QueryFriendsInvite(c *gin.Context) {
	panic("not implemented")
}

func (friendInviteApi *FriendInviteApi) InviteFriend(c *gin.Context) {
	caller.Call(friendInviteApi.InviteFriendService.InviteFriend, c, new(service.InviteFriendCmd))
}

func (friendInviteApi *FriendInviteApi) ReceiveInvite(c *gin.Context) {
	caller.Call(friendInviteApi.InviteFriendService.ReceiveInvite, c, new(service.ReceiveInviteCmd))
}

func (friendInviteApi *FriendInviteApi) RejectInvite(c *gin.Context) {

}
