package service

import (
	"douyin/models"
	"errors"
)

type UserInfoFlow struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	UserID        int64  `json:"-"`              // 发起用户id, 不进行序列化
	QueryUserID   int64  `json:"id"`             // 要查询的用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Username      string `json:"name"`           // 用户名称
}

func NewUserInfoFlow(userID int64, queryUserID int64) *UserInfoFlow {
	return &UserInfoFlow{UserID: userID, QueryUserID: queryUserID}
}

// Do 获取UserID相关信息
func (user *UserInfoFlow) Do() error {
	userInfoDAO := &models.UserInfo{UserID: user.UserID}
	if err := userInfoDAO.QueryByUserID(); err != nil {
		return errors.New("系统错误")
	}
	userFollowDAO := &models.UserFollow{UserFollowID: user.UserID, UserFollowedID: user.QueryUserID}
	user.IsFollow = userFollowDAO.UserRecordExist()
	user.Username = userInfoDAO.Username
	user.FollowCount = userInfoDAO.FollowCount
	user.FollowerCount = userInfoDAO.FollowerCount
	return nil
}
