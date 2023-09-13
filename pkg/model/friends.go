package model

import "time"

type FriendStatus int

const (
	NORMAL FriendStatus = iota // 接受邀请
	DELETE                     // 删除好友
)

type FriendInviteStatus int

const (
	INVITE  FriendInviteStatus = iota // 发起邀请
	REJECT                            // 拒绝邀请
	RECEIVE                           // 删除好友
)

type Friend struct {
	Id        int64        `mysql:"id" json:"id"`
	UserId    int64        `mysql:"user_id" json:"user_id"`
	FriendId  int64        `mysql:"friend_id" json:"friend_id"`
	Remark    string       `mysql:"remark" json:"remark"`
	Status    FriendStatus `mysql:"status" json:"status"`
	GmtCreate time.Time    `mysql:"gmt_create" json:"gmt_create"`
	GmtUpdate time.Time    `mysql:"gmt_update" json:"gmt_update"`
}

type FriendInvite struct {
	Id        int64              `mysql:"id" json:"id"`
	UserId    int64              `mysql:"user_id" json:"user_id"`
	FriendId  int64              `mysql:"friend_id" json:"friend_id"`
	Status    FriendInviteStatus `mysql:"status" json:"status"`
	Extra     string             `mysql:"extra" json:"extra"`
	GmtCreate time.Time          `mysql:"gmt_create" json:"gmt_create"`
	GmtUpdate time.Time          `mysql:"gmt_update" json:"gmt_update"`
}

func NewFriendInvite(userId int64, friendId int64) *FriendInvite {
	return &FriendInvite{
		UserId:    userId,
		FriendId:  friendId,
		Status:    INVITE,
		Extra:     "",
		GmtCreate: time.Now(),
		GmtUpdate: time.Now(),
	}
}

func (friend *FriendInvite) receiveInvite() {
	friend.Status = RECEIVE
	friend.GmtUpdate = time.Now()
}

func (friend *FriendInvite) rejectFriend() {
	friend.Status = REJECT
	friend.GmtUpdate = time.Now()
}

func (friend *Friend) deleteFriend() {
	friend.Status = DELETE
	friend.GmtUpdate = time.Now()
}
