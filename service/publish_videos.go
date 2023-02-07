package service

import (
	"douyin/models"
	"errors"
	"fmt"
)

type VideoFlow struct {
	UserID   int64
	Title    string
	VideoUrl string
	CoverUrl string
}

// GenerateFileName 生成用户投稿视频的服务端存储文件名，由user_id和其发布的视频数量拼接而成
func GenerateFileName(userID int64) (string, error) {
	var count int64
	videoDAO := &models.Video{UserID: userID}
	if err := videoDAO.QueryVideosCountByUserID(&count); err != nil {
		return "", errors.New("生成文件名错误")
	}
	return fmt.Sprintf("%d-%d", userID, count), nil
}

func PublishVideo(userID int64, title string, videoUrl string, coverUrl string) error {
	return NewVideoFlow(userID, title, videoUrl, coverUrl).Do()
}

func NewVideoFlow(userID int64, title string, videoUrl string, coverUrl string) *VideoFlow {
	return &VideoFlow{
		UserID:   userID,
		Title:    title,
		VideoUrl: videoUrl,
		CoverUrl: coverUrl,
	}
}

// Do 将VideoFlow中数据插入数据库中
func (video *VideoFlow) Do() error {
	videoDAO := &models.Video{
		UserID:   video.UserID,
		VideoUrl: video.VideoUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
	if err := videoDAO.SaveVideo(); err != nil {
		return err
	}
	return nil
}
