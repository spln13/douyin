package middlewares

import (
	"douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	}
}
