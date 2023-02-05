package models

import "time"

type Message struct {
	ID         int64
	Content    string
	UserFromID int64
	UserToID   int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
