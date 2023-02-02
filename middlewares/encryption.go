package middlewares

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// PasswordEncryptionMiddleware
// 使用SHA256算法对用户明文密码加密，向handles层发送加密后对密码进行后续处理
func PasswordEncryptionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		password := context.Query("password") // 获取password
		digest := sha256.New()                // 对密码加密
		digest.Write([]byte(password))
		passwordSHA := hex.EncodeToString(digest.Sum(nil))
		context.Set("password_sha256", passwordSHA) // 重写设置密码参数
		context.Next()                              // 放行
	}
}
