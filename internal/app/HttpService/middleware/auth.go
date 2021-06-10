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
	"xing-doraemon/internal/service/UserService"
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

var RedirectParam string = "return_url"

func LoginAuth(loginURL string) *Auth {
	return &Auth{
		Authenticator: func(ctx *gin.Context) error {
			u := UserService.GetUser(ctx)
			if !u.IsLogin() {
				return errors.New("no session")
			}
			err := UserService.UserSession.Save(ctx, u)
			if err != nil {
				return errors.New(fmt.Sprintf("update session err: %s", err.Error()))
			}
			ctx.Set("user", u)
			return nil
		},
		Unauthorized: func(ctx *gin.Context) {
			ctx.JSON(http.StatusFound, gin.H{
				"code": 302,
				"msg":  fmt.Sprintf("%s", loginURL),
				"data": "",
			})
		},
	}
}
