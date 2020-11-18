/*
 * @Time : 2020/11/16 10:43
 * @Author : wangyl
 * @File : auth.go
 * @Software: GoLand
 */
package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xing-doraemon/pkg/Auth/JwtAuth"
)

const (
	UserName = "UserName"
)

type Auth struct {
	Authenticator func(ctx *gin.Context) error
	Unauthorized  func(ctx *gin.Context)
}

func (a Auth) Func() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := a.Authenticator(ctx); err != nil {
			a.Unauthorized(ctx)
			ctx.Abort()
			return
		}
	}
}

func setUserName(ctx *gin.Context, userName string) {
	ctx.Set(UserName, userName)
}

func GetUserName(ctx *gin.Context) (string, bool) {
	if val, ok := ctx.Get(UserName); ok {
		strVal, ok := val.(string)
		return strVal, ok
	} else {
		return "", false
	}
}

type RedirectType string

const (
	RedirectTypeHttp RedirectType = "http_redirect"
	RedirectTypeJson RedirectType = "json_redirect"
)

var RedirectParam string = "return_url"

func LoginAuth(loginURL string, redirectType RedirectType) *Auth {
	return &Auth{
		Authenticator: func(ctx *gin.Context) error {
			token := ctx.Request.Header.Get("Authorization")
			if len(token) == 0 {
				return errors.New("no Authorization")
			}
			claims, err := JwtAuth.ParseToken(token)
			if err != nil {
				return err
			} else {
				setUserName(ctx, claims.Username)
				return nil
			}
		},
		Unauthorized: func(ctx *gin.Context) {
			if redirectType == RedirectTypeHttp {
				ctx.Redirect(http.StatusFound, loginURL)
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 302,
					"msg":  fmt.Sprintf("%s", loginURL),
					"data": "",
				})
			}
		},
	}
}
