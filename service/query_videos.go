package service

import (
	"douyin/models"
	"errors"
)

type ResponseModel struct {
	StatusCode  int64            `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg   *string          `json:"status_msg"`  // 返回状态描述
	VideoList   []QueryVideoFlow `json:"video_list"`  // 用户发布的视频列表
	QueryUserID int64            `json:"-"`           // 要查询的用户id
	UserID      int64            `json:"-"`           // 发起查询的用户id
}

type QueryVideoFlow struct {
	Author QueryUserFlow `json:"author"` // 视频作者信息
	//CommentCount  int64         `json:"comment_count"`  // 视频的评论总数
	//CoverURL      string        `json:"cover_url"`      // 视频封面地址
	//FavoriteCount int64         `json:"favorite_count"` // 视频的点赞总数
	//ID            int64         `json:"id"`             // 视频唯一标识
	IsFavorite bool `json:"is_favorite"` // true-已点赞，false-未点赞
	//PlayURL       string        `json:"play_url"`       // 视频播放地址
	//Title         string        `json:"title"`          // 视频标题
	VideoDAO *models.Video
}

type QueryUserFlow struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	UserID        int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Username      string `json:"name"`           // 用户名称
}

func (response *ResponseModel) Do() error {
	// 先根据QueryUserID查询作者信息
	userInfoDAO := models.UserInfo{UserID: response.QueryUserID}
	if err := userInfoDAO.QueryByUserID(); err != nil {
		return errors.New("获取用户信息错误")
	}
	author := &QueryUserFlow{
		FollowCount:   userInfoDAO.FollowCount,
		FollowerCount: userInfoDAO.FollowerCount,
		UserID:        userInfoDAO.UserID,
		IsFollow:      false,
		Username:      userInfoDAO.Username,
	}
	// 判断发起用户是否关注了查询用户
	userFollowDAO := &models.UserFollow{UserFollowID: response.UserID, UserFollowedID: response.QueryUserID}
	author.IsFollow = userFollowDAO.UserRecordExist()
	// 根据用户id查询该用户发布过的所有视频
	videoInfoList, err := models.QueryVideosByUserID(author.UserID)
	if err != nil {
		return errors.New("获取视频信息错误")
	}
	var videoFlowList []QueryVideoFlow
	for _, video := range *videoInfoList {
		oneVideoFlow := QueryVideoFlow{
			Author:     *author,
			IsFavorite: false,
			VideoDAO:   &video,
		}
		// 判断发起查询用户是否点赞video视频
		oneVideoFlow.IsFavorite = models.NewUserLikeDAO(response.UserID, video.ID).IsFavoriteRecordExists()
		videoFlowList = append(videoFlowList, oneVideoFlow)
	}
	response.VideoList = videoFlowList
	return nil
}
