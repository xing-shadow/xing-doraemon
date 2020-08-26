/*
@Time : 2020/7/23 16:37
@Author : wangyl
@File : HookHandle.go
@Software: GoLand
*/
package Hook

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"iPublic/LoggerModular"
)

func HookHandle(ctx *gin.Context) {
	logger := LoggerModular.GetLogger()
	var hookInfo Hook
	if err := ctx.BindJSON(&hookInfo); err != nil {
		logger.Error(err)
	}
	logger.WithField("RuleID", hookInfo.RuleID).Info(hookInfo)
	ctx.JSON(http.StatusOK, nil)
}
