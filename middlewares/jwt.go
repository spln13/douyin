package middlewares

import (
	"douyin/models"
	"douyin/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 生成token
var signature = "douyinSignature"

type Claims struct {
	ID                 int64
	jwt.StandardClaims // jwt中标准格式,主要是设置token的过期时间
}

// GenerateToken
// 调用库的NewWithClaims(加密方式,载荷).SignedString(签名) 生成token
func GenerateToken(u *service.UserRegisterLoginFlow) string {
	nowTime := time.Now()
	expirationTime := nowTime.Add(7 * 24 * time.Hour) // 过期时间
	issuer := "linan"
	claims := Claims{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signature))
	if err != nil {
		log.Println(err)
	}
	u.Token = token
	return token
}

// JWTMiddleware
// 如果token无效或者是过期则终止访问，否则获取token对应的UserID调用Handles
func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Query("token")
		if token == "" {
			context.JSON(http.StatusOK, models.CommonResponseBody{
				StatusCode:    1,
				StatusMessage: "token无效",
			})
			context.Abort()
			return
		}
		user := models.QueryUserByToken(token)
		if user.ID == 0 {
			context.JSON(http.StatusOK, models.CommonResponseBody{
				StatusCode:    2,
				StatusMessage: "用户不存在",
			})
			context.Abort()
			return
		}
		if time.Now().Unix() > user.TokenExpirationTime.Unix() {
			context.JSON(http.StatusOK, models.CommonResponseBody{
				StatusCode:    3,
				StatusMessage: "token过期",
			})
			context.Abort()
			return
		}
		context.Set("user_id", user.ID)
		context.Next()
	}
}
