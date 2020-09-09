/*
@Time : 2020/9/7 19:46
@Author : wangyl
@File : configsHandler.go
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

type ConfigHandler struct {
}

func (*ConfigHandler) GetAll(ctx *gin.Context) {
	idc := ctx.Query("idc")
	var config = &model.Configs{}
	data := config.GetAllConfig(Invoker.GetDB(), idc)
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func (*ConfigHandler) AddConfig(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var config = &model.Configs{}
	var resp common.Res
	err := ctx.BindJSON(&config)
	if err != nil {
		logger.Errorf("Unmarshal prom error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		err = config.AddConfig(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, config)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*ConfigHandler) UpdateConfig(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var config = &model.Configs{}
	var resp common.Res
	configId := ctx.Param("id")
	id, _ := strconv.ParseInt(configId, 10, 64)
	err := ctx.BindJSON(&config)
	if err != nil {
		logger.Errorf("Unmarshal prom error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		config.Id = id
		err = config.UpdateConfig(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Info("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, config)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*ConfigHandler) DeleteConfig(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var config = &model.Configs{}
	var resp common.Res
	configId := ctx.Param("id")
	err := config.DeleteConfig(Invoker.GetDB(), configId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, configId)
	ctx.JSON(http.StatusOK, resp)
}
