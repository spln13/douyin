package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandle(context *gin.Context) {
	username := context.Query("username")
	password := context.MustGet("password_sha256")
	fmt.Println(username, password)
}

func UserLoginHandle(context *gin.Context) {

}

func GetUserInfoHandle(context *gin.Context) {

}
