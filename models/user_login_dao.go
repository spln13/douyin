package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserLogin struct {
	ID                  int64  `gorm:"primary_key"`
	Username            string `gorm:"primary_key"`
	Password            string
	Token               string
	CreateTime          time.Time
	UpdateTime          time.Time
	TokenExpirationTime time.Time
}

// CheckUsernameUnique
// 检查该用户名是否唯一
func (u *UserLogin) CheckUsernameUnique() error {
	var user UserInfo
	err := GetDB().Where("username = ?", u.Username).Find(&user).Error
	if err != nil {
		log.Println(err)
	}
	if user.ID == 0 {
		return nil
	}
	return errors.New("用户名已存在")
}

// SaveUser
// 向数据库中插入新用户，若出错则回滚
func (u *UserLogin) SaveUser() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		// 进行插入事务
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		userInfo := &UserInfo{
			UserID:        u.ID,
			FollowCount:   0,
			FollowerCount: 0,
			UpdateTime:    time.Now(),
		}
		if err := tx.Create(userInfo).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("注册错误，请重试")
	}
	return nil
}

// UpdateUserToken
// 根据主键更新数据库中用户数据
func (u *UserLogin) UpdateUserToken() error {
	u.UpdateTime = time.Now()
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		if err := GetDB().Model(&u).Select("token", "update_time", "token_expiration_time").Updates(u).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// QueryByUsername
// 查询用户名对应的用户
func QueryByUsername(username string) *UserLogin {
	var user UserLogin
	err := GetDB().Where("username = ?", username).Select("id", "password").Find(&user).Error
	if err != nil {
		log.Println(err)
	}
	return &user
}
