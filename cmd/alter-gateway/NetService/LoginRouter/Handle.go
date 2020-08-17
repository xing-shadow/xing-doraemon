/*
@Time : 2020/7/20 11:26
@Author : wangyl
@File : GetMethodHandle.go
@Software: GoLand
*/
package LoginRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"xing-doraemon/cmd/alter-gateway/Common"
)

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

func MethodHandle(ctx *gin.Context) {
	//TODO 验证Session
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "OK"
	ctx.JSON(http.StatusOK,resp)
}

func UsernameHandle(ctx *gin.Context)  {
	//TODO 验证Session
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "OK"
	resp.Data = "xing"
	ctx.JSON(http.StatusOK,resp)
}

func localHandle(ctx *gin.Context)  {
	//TODO 验证User
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "Success"
	ctx.JSON(http.StatusOK,resp)
}

