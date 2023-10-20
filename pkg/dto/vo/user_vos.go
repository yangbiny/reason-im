package vo

import "time"

type UserVo struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	GmtCreate string `json:"gmt_create"`
}

type UserRelationVo struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	IsFriend  bool      `json:"is_friend"`
	GmtCreate time.Time `json:"gmt_create"`
}
