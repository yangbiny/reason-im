package api

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/application"
)

type FriendApi struct {
	FriendService *application.FriendService
}

func NewFriendApi(friendService *application.FriendService) *FriendApi {
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
