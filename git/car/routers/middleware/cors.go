// @Author zhangjiaozhu 2023/11/28 13:28:00
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	config.AllowOrigins = []string{"http://localhost:5000","http://127.0.0.1:*", "http://localhost:5000"}
	config.AllowCredentials = true
	return cors.New(config)
}


