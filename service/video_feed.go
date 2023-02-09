package service

import (
	"douyin/models"
	"errors"
	"log"
	"time"
)

var (
	videoLimit = 30
	staticPath = "192.168.2.105:8080/static/"
)

type FeedResponseFlow struct {
	NextTime   int64           `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	VideoList  []FeedVideoFlow `json:"video_list"`  // 视频列表
	LatestTime time.Time       `json:"-"`
}

type FeedVideoFlow struct {
	Author        FeedUserFlow `json:"author"`         // 视频作者信息
	CommentCount  int64        `json:"comment_count"`  // 视频的评论总数
	CoverURL      string       `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64        `json:"favorite_count"` // 视频的点赞总数
	ID            int64        `json:"id"`             // 视频唯一标识
	IsFavorite    bool         `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string       `json:"play_url"`       // 视频播放地址
	Title         string       `json:"title"`          // 视频标题
}

type FeedUserFlow struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

func NewFeedResponseFlow(latestTime time.Time) *FeedResponseFlow {
	return &FeedResponseFlow{LatestTime: latestTime}
}

func (feedResponseFlow *FeedResponseFlow) Do() error {
	var videoListDAO []*models.Video
	if err := models.QueryVideo(videoLimit, feedResponseFlow.LatestTime, &videoListDAO); err != nil {
		log.Println(err.Error())
		return errors.New("获取视频错误")
	}
	var videoListFlow []FeedVideoFlow
	for _, videoDAO := range videoListDAO {
		authorID := videoDAO.UserID
		userDAO := &models.UserInfo{UserID: authorID}
		if err := userDAO.QueryByUserID(); err != nil {
			return errors.New("获取作者信息错误")
		}
		authorFeedFlow := &FeedUserFlow{
			FollowCount:   userDAO.FollowCount,
			FollowerCount: userDAO.FollowerCount,
			ID:            authorID,
			IsFollow:      false, // 无登录状态
			Name:          userDAO.Username,
		}
		videoFeedFlow := FeedVideoFlow{
			Author:        *authorFeedFlow,
			CommentCount:  videoDAO.CommentCount,
			CoverURL:      staticPath + videoDAO.CoverUrl,
			PlayURL:       staticPath + videoDAO.VideoUrl,
			FavoriteCount: videoDAO.FavoriteCount,
			ID:            videoDAO.ID,
			IsFavorite:    false, // 无登陆状态
			Title:         videoDAO.Title,
		}
		videoListFlow = append(videoListFlow, videoFeedFlow)
	}
	feedResponseFlow.VideoList = videoListFlow
	return nil
}
