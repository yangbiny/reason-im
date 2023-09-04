package model

import "time"

type Users struct {
	Id        int64     `mysql:"id" json:"id"`
	Name      string    `mysql:"name" json:"name"`
	GmtCreate time.Time `mysql:"gmt_create" json:"gmt_create"`
	GmtUpdate time.Time `mysql:"gmt_update" json:"gmt_update"`
}
