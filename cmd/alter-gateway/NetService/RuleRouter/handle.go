/*
@Time : 2020/7/20 11:48
@Author : wangyl
@File : handle.go
@Software: GoLand
*/
package RuleRouter

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"iPublic/LoggerModular"

	"xing-doraemon/cmd/alter-gateway/model"
)

type Rule struct {
	Id    int64  `json:"id"`
	Expr  string `json:"expr"`
	Op    string `json:"op"`
	Value string `json:"value"`
	For   string `json:"for"`
	//Labels      map[string]string `json:"labels"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PromId      int64  `json:"prom_id"`
	PlanId      int64  `json:"plan_id"`
}

func handleGet(ctx *gin.Context)  {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			LoggerModular.GetLogger().Error("Panic in SendAllRules:%v\n%s", e, buf)
		}
	}()
	prom := ctx.Query("prom")
	id := ctx.Query("id")
	var Receiver *model.Rules
}
