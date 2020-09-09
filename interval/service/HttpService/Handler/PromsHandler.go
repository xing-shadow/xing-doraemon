/*
@Time : 2020/9/7 19:46
@Author : wangyl
@File : PromsHandler.go
@Software: GoLand
*/
package Handler

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type PromsHandler struct {
}

func (*PromsHandler) GetAllProms(ctx *gin.Context) {
	var Receiver = &model.Proms{}
	proms := Receiver.GetAllProms(Invoker.GetDB())
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: proms,
	})
}

func (*PromsHandler) AddProm(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var prom model.Proms
	var resp common.Res
	err := ctx.BindJSON(&prom)
	if err != nil {
		logger.Errorf("Unmarshal prom error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		err = prom.AddProms(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, prom)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*PromsHandler) UpdateProm(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var prom model.Proms
	var resp common.Res
	promId := ctx.Param("id")
	id, _ := strconv.ParseInt(promId, 10, 64)
	err := ctx.BindJSON(&prom)
	if err == nil {
		prom.Id = id
		err = prom.UpdateProms(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, prom)
	} else {
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*PromsHandler) DeleteProm(ctx *gin.Context) {
	logger := gobal.GetLogger()
	promId := ctx.Param("id")
	var Receiver = &model.Proms{}
	var resp common.Res
	err := Receiver.DeleteProms(Invoker.GetDB(), promId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, promId)
	ctx.JSON(http.StatusOK, resp)
}
