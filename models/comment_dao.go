package models

import "time"

type Comment struct {
	ID        int64
	UserID    int64
	VideoID   int64
	Content   string
	Likes     int64
	IsDel     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
