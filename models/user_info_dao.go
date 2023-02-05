package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID            int64 `gorm:"primary_key"`
	UserID        int64
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (user *UserInfo) IncreaseFollowCount() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("user_id = ?", user.UserID).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func (user *UserInfo) IncreaseFollowerCount() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("user_id = ?", user.UserID).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func (user *UserInfo) DecreaseFollowCount() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("user_id = ?", user.UserID).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func (user *UserInfo) DecreaseFollowerCount() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("user_id = ?", user.UserID).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}
