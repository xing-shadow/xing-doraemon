/*
@Time : 2020/9/7 15:21
@Author : wangyl
@File : LogoutHandler.go
@Software: GoLand
*/
package Handler

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"xing-doraemon/pkg/common"
)

type LogoutHandler struct {
}

func (l *LogoutHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	resp := &common.Res{
		Code: 0,
		Msg:  "Success",
	}
	ctx.JSON(http.StatusOK, resp)
}
