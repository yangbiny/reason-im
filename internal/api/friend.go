package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/pkg/service"
)

type FriendApi struct {
	FriendService *service.FriendService
}

func NewFriendApi(friendService *service.FriendService) *FriendApi {
	return &FriendApi{
		FriendService: friendService,
	}
}

func (friendApi *FriendApi) QueryFriends(c *gin.Context) {
	panic("")
}

func (friendApi *FriendApi) DeleteFriend(c *gin.Context) {
	panic("")
}
