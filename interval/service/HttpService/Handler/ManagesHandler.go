/*
@Time : 2020/9/7 19:46
@Author : wangyl
@File : ManagesHandler.go
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

type ManagerHandler struct {
}

func (*ManagerHandler) GetAll(ctx *gin.Context) {
	var manage = &model.Manages{}
	res := manage.GetAllManage(Invoker.GetDB())
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: res,
	})
}

func (*ManagerHandler) AddManage(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var manage model.Manages
	var resp common.Res
	err := ctx.BindJSON(&manage)
	if err != nil {
		logger.Errorf("Unmarshal prom error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		err = manage.AddManage(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, manage)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*ManagerHandler) UpdateManage(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var manage model.Manages
	var resp common.Res
	manageId := ctx.Param("id")
	id, _ := strconv.ParseInt(manageId, 10, 64)
	err := ctx.BindJSON(&manage)
	if err != nil {
		logger.Error("Unmarshal prom error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		manage.Id = id
		err = manage.UpdateManage(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, manage)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*ManagerHandler) DeleteManage(ctx *gin.Context) {
	logger := gobal.GetLogger()
	manageId := ctx.Param("id")
	var manage = &model.Manages{}
	var resp common.Res
	err := manage.DeleteManage(Invoker.GetDB(), manageId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, manageId)
	ctx.JSON(http.StatusOK, resp)
}
