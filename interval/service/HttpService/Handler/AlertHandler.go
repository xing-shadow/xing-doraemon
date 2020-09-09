/*
@Time : 2020/9/7 17:28
@Author : wangyl
@File : AlertHandler.go
@Software: GoLand
*/
package Handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type AlertHandler struct {
}

func (*AlertHandler) GetAlerts(ctx *gin.Context) {
	pageNo, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.Query("pagesize"), 10, 64)
	timeStart := ctx.Query("timestart")
	timeEnd := ctx.Query("timeend")
	status := ctx.Query("status")
	summary := ctx.Query("summary")
	if pageNo == 0 && pageSize == 0 {
		pageNo = 1
		pageSize = 10
	}
	var Receiver = &model.Alerts{}
	alerts := Receiver.GetAlerts(Invoker.GetDB(), pageNo, pageSize, timeStart, timeEnd, status, summary)
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: alerts,
	})
}

func (*AlertHandler) ShowAlerts(ctx *gin.Context) {
	ruleId := ctx.Param(":ruleid")
	start := ctx.Query("start")
	pageNo, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.Query("pagesize"), 10, 64)
	var Receiver = model.Alerts{}
	alerts := Receiver.ShowAlerts(Invoker.GetDB(), ruleId, start, pageNo, pageSize)
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: alerts,
	})
}

func (*AlertHandler) ClassifyAlerts(ctx *gin.Context) {
	var Receiver = &model.Alerts{}
	alerts := Receiver.ClassifyAlerts(Invoker.GetDB())
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: alerts,
	})
}

func (*AlertHandler) Confirm(ctx *gin.Context) {
	var confirmList common.Confirm
	var resp common.Res
	err := ctx.BindWith(&confirmList, binding.JSON)
	if err == nil {
		var Receiver = &model.Alerts{}
		err = Receiver.ConfirmAll(Invoker.GetDB(), &confirmList)
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		gobal.GetLogger().Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, confirmList)
	} else {
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*AlertHandler) HandleAlerts(ctx *gin.Context) {
	var alerts common.Alerts
	var resp common.Res
	err := ctx.BindWith(&alerts, binding.JSON)
	gobal.GetLogger().Infof("%v\n", alerts)
	if err != nil {
		gobal.GetLogger().Errorf("Unmarshal error:%s", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		var Receiver = model.Alerts{}
		Receiver.AlertsHandler(Invoker.GetDB(), &alerts)
	}
	ctx.JSON(http.StatusOK, resp)
}
