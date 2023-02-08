package service

import (
	"douyin/models"
)

type FavoriteVideoFlow struct {
	UserID     int64
	VideoID    int64
	ActionType int
}

func NewFavoriteVideoFlow(userID int64, videoID int64, actionType int) *FavoriteVideoFlow {
	return &FavoriteVideoFlow{
		UserID:     userID,
		VideoID:    videoID,
		ActionType: actionType,
	}
}

func FavoriteAction(userID int64, videoID int64, actionType int) error {
	return NewFavoriteVideoFlow(userID, videoID, actionType).Do()
}

// Do 根据ActionType向数据库中插入或者删除点赞视频记录
func (video *FavoriteVideoFlow) Do() error {
	videoLikeDAO := models.NewUserLikeDAO(video.UserID, video.VideoID)
	if video.ActionType == 1 {
		if err := videoLikeDAO.InsertFavoriteRecord(); err != nil {
			return err
		}
	}
	if video.ActionType == 2 {
		if err := videoLikeDAO.DeleteFavoriteRecord(); err != nil {
			return err
		}
	}
	return nil
}
