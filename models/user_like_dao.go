package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserLike struct {
	ID        int64
	UserID    int64
	VideoID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserLikeDAO(userID int64, videoID int64) *UserLike {
	return &UserLike{UserID: userID, VideoID: videoID}
}

func (userLikeDAO *UserLike) InsertFavoriteRecord() error {
	if exists := userLikeDAO.IsFavoriteRecordExists(); exists {
		return errors.New("您已经点过赞了")
	}
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(userLikeDAO).Error
	})
	if err != nil {
		return errors.New("点赞错误")
	}
	return nil
}

func (userLikeDAO *UserLike) DeleteFavoriteRecord() error {
	if exists := userLikeDAO.IsFavoriteRecordExists(); !exists {
		return errors.New("您未点过赞")
	}
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Model(&UserLike{}).Delete(userLikeDAO).Error // 可直接删除userLikeDAO因其主键ID已在IsFavoriteRecordExists()中赋值
	})
	if err != nil {
		return errors.New("取消点赞错误")
	}
	return nil
}

// IsFavoriteRecordExists 根据UserID和VideoID判断该用户是否对该视频点过赞
func (userLikeDAO *UserLike) IsFavoriteRecordExists() bool {
	err := GetDB().Where("user_id = ? and video_id = ?", userLikeDAO.UserID, userLikeDAO.VideoID).Select("id").Find(userLikeDAO).Error
	if err != nil {
		log.Println(err.Error())
	}
	if userLikeDAO.ID == 0 {
		return false
	}
	return true
}
