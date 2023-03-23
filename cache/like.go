package cache

import (
	"errors"
	"log"
	"strconv"
)

func LikeVideo(userID, videoID int64) error {
	err := rdb.SAdd(ctx, strconv.FormatInt(userID, 10), videoID).Err()
	if err != nil {
		log.Println(err)
		return errors.New("点赞视频失败")
	}
	return nil
}

func UnLikeVideo(userID, videoID int64) error {
	err := rdb.SRem(ctx, strconv.FormatInt(userID, 10), videoID)
	if err != nil {
		log.Println(err)
		return errors.New("取消点赞失败")
	}
	return nil
}
