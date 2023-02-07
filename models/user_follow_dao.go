package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserFollow struct {
	ID             int64
	UserFollowID   int64
	UserFollowedID int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// UserRecordExist 根据UserFollowID和UserFollowedID检索对应关注记录是否存在
func (user *UserFollow) UserRecordExist() bool {
	err := GetDB().Where("user_follow_id = ? and user_followed_id = ?", user.UserFollowID, user.UserFollowedID).Select("id").Find(&user).Error
	if err != nil {
		return false
	}
	if user.ID == 0 { // 不存在记录
		return false
	}
	return true
}

func (user *UserFollow) InsertUserFollow() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func (user *UserFollow) DeleteUserFollow() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}
