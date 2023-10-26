package vo

type GroupVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type InviteGroupVo struct {
	Id           int64 `json:"id"`
	InviteUserId int64 `json:"invite_user_id"`
	GroupId      int64 `json:"group_id"`
}
