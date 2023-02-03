package handlers

import (
	"douyin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	MaxUsernameLength = 30
)

type RegisterResponse struct {
	models.CommonResponseBody
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserRegisterHandle(context *gin.Context) {
	username := context.Query("username")
	rawData, exists := context.Get("password_sha256")
	password, ok := rawData.(string)
	if !exists || !ok { // 密码解析出错
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    1,
				StatusMessage: "密码解析错误",
			},
		})
		return
	}
	user := models.UserLogin{
		Username:   username,
		Password:   password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := user.CheckParamValid(MaxUsernameLength) // 检查用户名是否符合规范
	if err != nil {                                // 出错
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    2,
				StatusMessage: err.Error(),
			},
		})
		return
	}
	ok = user.CheckUsernameUnique() // 检查该用户名是否唯一
	if !ok {
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    3,
				StatusMessage: "用户名已经存在",
			},
		})
		return
	}
	user.GenerateToken() // 生成token
	if err != nil {
		log.Println(err)
	}
	err = user.SaveUser()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    4,
				StatusMessage: "请重试",
			},
		})
		return
	}
	// 注册成功
	context.JSON(http.StatusOK, RegisterResponse{
		CommonResponseBody: models.CommonResponseBody{
			StatusCode: 0,
		},
		UserID: user.ID,
		Token:  user.Token,
	})
}

func UserLoginHandle(context *gin.Context) {
	username := context.Query("username")
	rawData, exists := context.Get("password_sha256")
	password, ok := rawData.(string) // password为加密后
	if !exists || !ok {              // 密码解析出错
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    1,
				StatusMessage: "密码解析错误",
			},
		})
		return
	}
	user := models.QueryByUsername(username)
	if user.Password != password {
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    1,
				StatusMessage: "密码错误",
			},
		})
		return
	}
	user.GenerateToken() // 更新token

}

func GetUserInfoHandle(context *gin.Context) {

}
