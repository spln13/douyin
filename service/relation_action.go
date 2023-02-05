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
		if err := action.Follow(); err != nil {
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
	if exists := userFollowDAO.UserRecordExist(); exists {
		return errors.New("您已经关注过该用户了 ")
	}
	if err := userFollowDAO.InsertUserFollow(); err != nil {
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
	return nil
}
