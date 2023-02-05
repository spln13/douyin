package service

import (
	"douyin/models"
	"errors"
)

type RelationActionFlow struct {
	UserID     int64
	ToUserID   int64
	ActionType int
}

func NewRelationActionFlow(userId int64, toUserId int64, actionType int) *RelationActionFlow {
	return &RelationActionFlow{
		UserID:     userId,
		ToUserID:   toUserId,
		ActionType: actionType,
	}
}

// Do
// 完成user_id向to_user_id关注或者取消关注
func (action *RelationActionFlow) Do() error {
	if action.ActionType == 1 { // 关注操作
		if err := action.Follow(); err != nil {
			return err
		}
	} else if action.ActionType == 2 { // 取消关注操作
		if err := action.Unfollow(); err != nil {
			return err
		}
	} else {
		if err := errors.New("操作解析错误"); err != nil {
			return err
		}
	}
	return nil
}

func (action *RelationActionFlow) Follow() error {
	userFollowDAO := &models.UserFollow{UserFollowID: action.UserID, UserFollowedID: action.ToUserID}
	if exists := userFollowDAO.UserRecordExist(); exists { // 查询用户关注记录是否存在
		return errors.New("您已经关注过该用户了 ")
	}
	if err := userFollowDAO.InsertUserFollow(); err != nil { // 插入用户关注记录
		return err
	}
	// 更改user_info表中follow_count和follower_count
	userInfoDao := &models.UserInfo{UserID: userFollowDAO.UserFollowID}
	userToInfoDao := &models.UserInfo{UserID: userFollowDAO.UserFollowedID}
	if err := userInfoDao.IncreaseFollowCount(); err != nil {
		return err
	}
	if err := userToInfoDao.IncreaseFollowerCount(); err != nil {
		return err
	}
	return nil
}

func (action *RelationActionFlow) Unfollow() error {
	userFollowDAO := &models.UserFollow{UserFollowID: action.UserID, UserFollowedID: action.ToUserID}
	if exists := userFollowDAO.UserRecordExist(); !exists {
		return errors.New("您尚未关注该用户")
	}
	if err := userFollowDAO.DeleteUserFollow(); err != nil {
		return err
	}
	userInfoDao := &models.UserInfo{UserID: userFollowDAO.UserFollowID}
	userToInfoDao := &models.UserInfo{UserID: userFollowDAO.UserFollowedID}
	if err := userInfoDao.DecreaseFollowCount(); err != nil {
		return err
	}
	if err := userToInfoDao.DecreaseFollowerCount(); err != nil {
		return err
	}
	return nil
}
