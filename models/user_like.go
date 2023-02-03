package models

import "time"

type UserLike struct {
	ID         int64
	UserID     int64
	VideoID    int64
	IsDel      bool
	CreateTime time.Time
}
