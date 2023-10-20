package vo

import "time"

type UserInviteVo struct {
	Id             int64     `json:"id"`
	InviteUserId   int64     `json:"invite_user_id"`
	InviteUserName string    `json:"invite_user_name"`
	InviteStatus   int       `json:"invite_status"`
	Extra          string    `json:"extra"`
	GmtCreate      time.Time `json:"gmt_create"`
}
