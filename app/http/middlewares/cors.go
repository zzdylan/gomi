// Package middlewares Gin 中间件
package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cors 允许跨域请求中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的来源，* 表示允许任何来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的HTTP方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// 设置是否允许携带认证信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理OPTIONS预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
