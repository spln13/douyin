package models

import "time"

type Video struct {
	ID         int64
	VideoUrl   string
	CoverUrl   string
	Title      string
	IsDel      bool
	CreateTime time.Time
}
