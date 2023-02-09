package handlers

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func FeedHandle(context *gin.Context) {
	// 默认无登陆状态
	feedResponse := service.NewFeedResponseFlow(time.Now())
	if err := feedResponse.Do(); err != nil {
		feedResponse.StatusMsg = err.Error()
		context.JSON(http.StatusOK, models.CommonResponseBody{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
	}
	feedResponse.StatusCode = 0
	log.Println(feedResponse)
	context.JSON(http.StatusOK, feedResponse)
}
