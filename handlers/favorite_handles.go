package handlers

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteActionHandle(context *gin.Context) {
	userID, ok := context.MustGet("user_id").(int64)
	if !ok {
		ErrorResponse(context, "token解析错误")
		return
	}
	videoIDStr := context.Query("video_id")
	videoID, _ := strconv.ParseInt(videoIDStr, 10, 64)
	actionTypeStr := context.Query("action_type")
	actionType, _ := strconv.Atoi(actionTypeStr) // 1-点赞，2-取消点赞
	if err := service.FavoriteAction(userID, videoID, actionType); err != nil {
		ErrorResponse(context, err.Error())
	}
	// 操作成功
	context.JSON(http.StatusOK, models.CommonResponseBody{
		StatusCode: 0,
	})

}
func GetFavoriteListHandle(context *gin.Context) {

}
