/*
@Time : 2020/7/20 10:25
@Author : wangyl
@File : middlerware.go
@Software: GoLand
*/
package AlterGateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Headers", "*,content-time")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

func FilterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if (len(ctx.Request.RequestURI) >= 13 && ctx.Request.RequestURI[:13] == "/api/v1/proms" && ctx.Request.Method == "GET" || len(ctx.Request.RequestURI) >= 13 && ctx.Request.RequestURI[:13] == "/api/v1/rules" && ctx.Request.Method == "GET" || len(ctx.Request.RequestURI) >= 14 && ctx.Request.RequestURI[:14] == "/api/v1/alerts" && ctx.Request.Method == "POST") && ctx.Input.Header("Token") == "96smhbNpRguoJOCEKNrMqQ" {
			return
		}
		if len(ctx.Request.RequestURI) >= 14 && ctx.Request.RequestURI[:14] == "/api/v1/logout" {
			return
		}
	}
}
