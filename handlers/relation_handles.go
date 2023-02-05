package handlers

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationActionHandle(context *gin.Context) {
	toUserIdStr := context.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(toUserIdStr, 10, 64)
	actionTypeStr := context.Query("action_type")
	actionType, _ := strconv.Atoi(actionTypeStr)
	userId, ok := context.MustGet("user_id").(int64)
	if !ok {
		context.JSON(http.StatusOK, models.CommonResponseBody{
			StatusCode:    1,
			StatusMessage: "token解析错误",
		})
		return
	}
	relationActionFlow := service.NewRelationActionFlow(userId, toUserId, actionType)
	err := relationActionFlow.Do()
	if err != nil {
		context.JSON(http.StatusOK, models.CommonResponseBody{
			StatusCode:    2,
			StatusMessage: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, models.CommonResponseBody{
		StatusCode:    0,
		StatusMessage: "操作成功",
	})
	return
}

func GetFollowListHandle(context *gin.Context) {

}

func GetFollowerListHandle(context *gin.Context) {

}

func GetFriendListHandle(context *gin.Context) {

}
