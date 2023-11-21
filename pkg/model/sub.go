package model

import "time"

type Sub struct {
	Id        int64
	Name      string
	OwnerId   int64
	GmtCreate time.Time
	GmtUpdate time.Time
}
