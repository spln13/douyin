package dao

import "time"

type UserInfo struct {
	ID            int64
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	CreateTime    time.Time
	UpdateTime    time.Time
}
