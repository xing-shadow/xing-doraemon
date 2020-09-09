/*
@Time : 2020/9/7 19:47
@Author : wangyl
@File : GroupHandler.go
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

type GroupHandler struct {
}

func (*GroupHandler) GetAllGroup(ctx *gin.Context) {
	var Receiver = &model.Groups{}
	groups := Receiver.GetAll(Invoker.GetDB())
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: groups,
	})
}

func (*GroupHandler) AddGroup(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var group model.Groups
	var resp common.Res
	err := ctx.BindJSON(&group)
	if err != nil {
		logger.Errorf("Unmarshal plan error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		err = group.AddGroup(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, group)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*GroupHandler) UpdateGroup(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var group model.Groups
	var resp common.Res
	groupId := ctx.Param("id")
	err := ctx.BindJSON(&group)
	if err == nil {
		id, _ := strconv.ParseInt(groupId, 10, 64)
		group.Id = id
		err = group.UpdateGroup(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, group)
	} else {
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*GroupHandler) DeleteGroup(ctx *gin.Context) {
	logger := gobal.GetLogger()
	groupId := ctx.Param("id")
	var Receiver = &model.Groups{}
	var resp common.Res
	err := Receiver.DeleteGroup(Invoker.GetDB(), groupId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, groupId)
	ctx.JSON(http.StatusOK, resp)
}
