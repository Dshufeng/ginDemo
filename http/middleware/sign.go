package middleware

import (
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
)


func SignMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		utils.VerifySign(context)
		context.Next()
	}
}
