package dao

import "time"

type UserFollow struct {
	ID             int64
	UserFollowID   int64
	UserFollowedID int64
	CreateTime     time.Time
	IsDel          bool
}
