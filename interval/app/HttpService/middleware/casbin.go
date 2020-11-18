/*
 * @Time : 2020/11/18 10:35
 * @Author : wangyl
 * @File : casbin.go
 * @Software: GoLand
 */
package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xing-doraemon/interval/service/CasbinService"
	"xing-doraemon/pkg/App/Resp"
)

type CasbinConfig struct {
	Enforcer *casbin.SyncedEnforcer
	Skipper  func(ctx *gin.Context) bool
}

func AllowPathPrefixSkipper(prefixes ...string) func(ctx *gin.Context) bool {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		for _, prefix := range prefixes {
			if strings.HasPrefix(path, prefix) {
				return true
			}
		}
		return false
	}
}

func CasbinMiddleware(config CasbinConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if config.Skipper(ctx) {
			return
		}
		user, ok := GetUserName(ctx)
		if !ok {
			ctx.JSON(http.StatusOK, Resp.Response{
				Code:    Resp.MsgError,
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		pass, err := config.CheckPermission(user, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			ctx.JSON(http.StatusOK, Resp.Response{
				Code:    Resp.MsgError,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}
		if pass {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, Resp.Response{
				Code:    Resp.MsgError,
				Message: "用户没有该接口权限",
			})
			ctx.Abort()
			return
		}
	}
}

func (c CasbinConfig) CheckPermission(userName string, api string, method string) (bool, error) {
	return CasbinService.Casbin.CheckPermission(userName, api, method)
}
