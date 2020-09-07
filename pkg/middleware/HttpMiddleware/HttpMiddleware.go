/*
@Time : 2020/9/7 14:00
@Author : wangyl
@File : HttpMiddleware.go
@Software: GoLand
*/
package HttpMiddleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"xing-doraemon/pkg/common"
)

func FilterUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		requestURI := req.RequestURI
		method := req.Method
		if ((method == "GET" && len(requestURI) >= 13 && (requestURI[:13] == "/api/v1/proms" || requestURI[:13] == "/api/v1/rules")) ||
			(method == "POST" && len(requestURI) >= 14 && requestURI[:14] == "/api/v1/alerts")) &&
			ctx.Request.Header.Get("Token") == "96smhbNpRguoJOCEKNrMqQ" {
			ctx.Next()
		}
		if len(requestURI) >= 14 && requestURI[:14] == "/api/v1/logout" {
			ctx.Next()
		}
		session := sessions.Default(ctx)
		username, _ := session.Get("username").(string)
		if username == "" && ctx.Request.RequestURI[:13] != "/api/v1/login" {
			ctx.JSON(http.StatusUnauthorized, common.Res{Code: -1, Msg: "Unauthorized"})
		}
	}
}
