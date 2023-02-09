package handlers

import (
	"douyin/models"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	videoIndexMap = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
)

func PublishVideoHandle(context *gin.Context) {
	userID, ok := context.MustGet("user_id").(int64)
	if !ok {
		ErrorResponse(context, "token解析错误")
		return
	}
	title := context.PostForm("title")
	log.Println("to here is ok")
	file, err := context.FormFile("data")
	if err != nil {
		ErrorResponse(context, "请求解析错误")
		return
	}
	suffix := filepath.Ext(file.Filename)    // 获取文件类型后缀
	if _, ok := videoIndexMap[suffix]; !ok { // 判断是否为视频格式
		ErrorResponse(context, "视频格式不支持")
		return
	}
	uniqueName, err := service.GenerateFileName(userID)
	if err != nil {
		ErrorResponse(context, err.Error())
		return
	}
	fileName := uniqueName + suffix
	savePath := filepath.Join("./static", fileName)
	log.Println("videoSavePath is:", savePath)
	if err = context.SaveUploadedFile(file, savePath); err != nil {
		ErrorResponse(context, "存储文件错误01")
		return
	}
	// 截取视频的一帧作为封面存储
	go func() {
		if err = util.SaveImageFromVideo(uniqueName, true); err != nil {
			ErrorResponse(context, "存储文件错误02")
			return
		}
	}()
	// 将文件信息存入数据库
	if err = service.PublishVideo(userID, title, fileName, uniqueName+".jpg"); err != nil {
		ErrorResponse(context, "存储文件错误03")
	}
	// 执行成功
	context.JSON(http.StatusOK, models.CommonResponseBody{
		StatusCode: 0,
	})
}

func GetPublishListHandle(context *gin.Context) {
	userID, ok := context.MustGet("user_id").(int64)
	if !ok {
		ErrorResponse(context, "token解析错误")
		return
	}
	queryUserIDStr := context.Query("user_id")
	queryUserID, _ := strconv.ParseInt(queryUserIDStr, 10, 64)
	response := service.ResponseModel{QueryUserID: queryUserID, UserID: userID}
	if err := response.Do(); err != nil {
		ErrorResponse(context, "token解析错误")
	}
	context.JSON(http.StatusOK, response)
}

func ErrorResponse(context *gin.Context, message string) {
	context.JSON(http.StatusOK, models.CommonResponseBody{
		StatusCode:    1,
		StatusMessage: message,
	})
}
