package api

import (
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
