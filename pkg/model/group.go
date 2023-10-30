package model

import "time"

type GroupType int
type GroupMemberRole int

const (
	// GROUP_NORMAL 1. 普通群组
	GROUP_NORMAL GroupType = iota
	// GROUP_TEMP 2. 临时群组
	GROUP_TEMP

	// GROUP_MEMBER_ROLE_NORMAL 1. 普通群组成员
	GROUP_MEMBER_ROLE_NORMAL GroupMemberRole = iota
	// GROUP_MEMBER_ROLE_ADMIN 2. 群组管理员
	GROUP_MEMBER_ROLE_ADMIN
	// GROUP_MEMBER_ROLE_OWNER 3. 群组所有者
	GROUP_MEMBER_ROLE_OWNER
)

type Group struct {
	Id          int64
	Name        string
	Description string
	// 群组类型
	GroupType      GroupType
	GroupMemberCnt int32
	GmtCreate      time.Time
	GmtUpdate      time.Time
}

type GroupSettings struct {
	Id int64
	// 群组id
	GroupId int64
}

type GroupMember struct {
	Id int64
	// 群组id
	GroupId int64
	// 用户id
	UserId int64
	// 群组昵称
	NickName string
	// 群组成员角色
	GroupMemberRole GroupMemberRole
	// 群组成员加入时间
	GmtCreate time.Time
	// 群组成员更新时间
	GmtUpdate time.Time
}

func (groupMember *GroupMember) CanInviteUser() bool {
	return groupMember.GroupMemberRole == GROUP_MEMBER_ROLE_ADMIN || groupMember.GroupMemberRole == GROUP_MEMBER_ROLE_OWNER
}
