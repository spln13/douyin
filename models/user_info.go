package models

import (
	"time"
)

type UserInfo struct {
	ID            int64 `gorm:"primary_key"`
	UserID        int64
	FollowCount   int64
	FollowerCount int64
	UpdateTime    time.Time
}
