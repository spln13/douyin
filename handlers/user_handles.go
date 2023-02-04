package handlers

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	UserRegisterFlow := service.NewUserRegisterLoginFlow(username, password)
	if err := UserRegisterFlow.DoRegister(); err != nil { // 出错
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    2,
				StatusMessage: err.Error(),
			},
		})
		return
	}

	// 注册成功
	context.JSON(http.StatusOK, RegisterResponse{
		CommonResponseBody: models.CommonResponseBody{
			StatusCode:    0,
			StatusMessage: "注册成功",
		},
		UserID: UserRegisterFlow.ID,
		Token:  UserRegisterFlow.Token,
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
	userLoginFlow := service.NewUserRegisterLoginFlow(username, password)
	if err := userLoginFlow.DoLogin(); err != nil {
		context.JSON(http.StatusOK, RegisterResponse{
			CommonResponseBody: models.CommonResponseBody{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		})
		return
	}
	// 登录成功，返回参数
	context.JSON(http.StatusOK, RegisterResponse{
		CommonResponseBody: models.CommonResponseBody{
			StatusCode:    0,
			StatusMessage: "登录成功",
		},
		UserID: userLoginFlow.ID,
		Token:  userLoginFlow.Token,
	})
}

func GetUserInfoHandle(context *gin.Context) {

}
