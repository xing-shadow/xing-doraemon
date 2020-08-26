/*
@Time : 2020/7/20 11:35
@Author : wangyl
@File : handle.go
@Software: GoLand
*/
package LogoutRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xing-doraemon/cmd/alter-gateway/Common"
)

func handle(ctx *gin.Context) {
	// TODO 删除session
	var resp Common.Res
	resp.Msg = "Success"
	resp.Code = 0
	ctx.JSON(http.StatusOK, resp)
}
