package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID        int64
	UserID    int64
	VideoUrl  string
	CoverUrl  string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// QueryVideosCountByUserID 查询user_id对应用户发布过多少条视频
func (video *Video) QueryVideosCountByUserID(count *int64) error {
	return GetDB().Model(&Video{}).Where("user_id = ?", video.UserID).Count(count).Error
}

func (video *Video) SaveVideo() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(video).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
