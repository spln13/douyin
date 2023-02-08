package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"-"`
	VideoUrl      string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	Title         string    `json:"title"`
	FavoriteCount int       `json:"favorite_count"`
	CommentCount  int       `json:"comment_count"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
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

func QueryVideosByUserID(userID int64) (*[]Video, error) {
	var videoList *[]Video
	err := GetDB().Model(&Video{}).Where("user_id = ?", userID).Omit("create_at", "update_at").Find(videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
