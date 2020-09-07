/*
@Time : 2020/9/7 14:47
@Author : wangyl
@File : LoginHandler.go
@Software: GoLand
*/
package Handler

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/auth/ldaputil"
	"xing-doraemon/pkg/common"
)

type LoginHandler struct {
}

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type UserInfo struct {
	Name         string `json:"name"`
	Display      string `json:"display"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"is_admin"`
	AccessToken  string `json:"access_token"`
	ClientID     string `json:"client_id"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func (l *LoginHandler) GetMethod(ctx *gin.Context) {
	session := sessions.Default(ctx)
	method := session.Get(session)
	if method == nil {
		ctx.JSON(http.StatusUnauthorized, &common.Res{
			Code: -1,
			Msg:  "Unauthorized",
		})
	} else {
		ctx.JSON(http.StatusOK, &common.Res{
			Code: 0,
			Msg:  "",
			Data: method.(string),
		})
	}
}

func (l *LoginHandler) Username(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")
	if username == nil {
		ctx.JSON(http.StatusUnauthorized, &common.Res{
			Code: -1,
			Msg:  "Unauthorized",
		})
	} else {
		ctx.JSON(http.StatusOK, &common.Res{
			Code: 0,
			Msg:  "",
			Data: username.(string),
		})
	}
}

func (l *LoginHandler) Local(ctx *gin.Context) {
	var auth common.AuthModel
	var res common.Res
	var User *model.Users

	if err := ctx.BindJSON(&auth); err != nil {
		ctx.JSON(http.StatusInternalServerError, &common.Res{
			Code: -1,
			Msg:  "bind param to json err:" + err.Error(),
		})
		return
	}
	userInfo, err := User.CheckUser(Invoker.GetDB(), auth)
	if err == nil {
		session := sessions.Default(ctx)
		session.Set("username", userInfo.Username)
		session.Set("method", "local")
		session.Save()
		res.Msg = "Success"
	} else {
		res.Code = -1
		res.Msg = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func (l *LoginHandler) Ldap(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var auth common.AuthModel
	var res common.Res

	if err := ctx.BindJSON(&auth); err != nil {
		ctx.JSON(http.StatusInternalServerError, &common.Res{
			Code: -1,
			Msg:  "bind param to json err:" + err.Error(),
		})
		return
	}
	if err := ldaputil.Authenticate(auth.Username, auth.Password); err != nil {
		logger.Infof("authenticate fail error: %v", err)
		res.Code = -1
		res.Msg = "Unauthorized"
	} else {
		session := sessions.Default(ctx)
		session.Set("username", auth.Username)
		session.Set("method", "ldap")
		session.Save()
		res.Msg = "Success"
	}
	ctx.JSON(http.StatusOK, res)
}
