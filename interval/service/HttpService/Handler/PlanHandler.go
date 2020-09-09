/*
@Time : 2020/9/7 19:45
@Author : wangyl
@File : PlanHandler.go
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

type PlanHandler struct {
}

func (*PlanHandler) GetAllPlans(ctx *gin.Context) {
	var Receiver = &model.Plans{}
	plans := Receiver.GetAllPlans(Invoker.GetDB())
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: plans,
	})
}

func (*PlanHandler) AddPlan(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var plan = &model.Plans{}
	var resp common.Res
	err := ctx.BindWith(&plan, binding.JSON)
	if err != nil {
		logger.Error("Unmarshal plan error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		err = plan.AddPlan(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, plan)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*PlanHandler) GetAllReceiver(ctx *gin.Context) {
	planId := ctx.Param(":planid")
	var Receiver = &model.Receivers{}
	receivers := Receiver.GetAllReceivers(Invoker.GetDB(), planId)
	ctx.JSON(http.StatusOK, &common.Res{
		Code: 0,
		Msg:  "",
		Data: receivers,
	})
}

func (*PlanHandler) AddReceiver(ctx *gin.Context) {
	logger := gobal.GetLogger()
	planId := ctx.Param(":planid")
	var receiver = &model.Receivers{}
	var resp common.Res
	err := ctx.BindWith(&receiver, binding.JSON)
	if err != nil {
		logger.Errorf("Unmarshal rule error:%v", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		if receiver.Expression != "" {
			root, err := common.BuildTree(receiver.Expression)
			if err != nil {
				resp.Code = 1
				resp.Msg = err.Error()
			} else {
				ReversePolishNotation := common.Converse2ReversePolishNotation(root)
				receiver.ReversePolishNotation = ReversePolishNotation
				id, _ := strconv.ParseInt(planId, 10, 64)
				receiver.Plan = &model.Plans{Id: id}
				err = receiver.AddReceiver(Invoker.GetDB())
				if err != nil {
					resp.Code = 1
					resp.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(planId, 10, 64)
			receiver.Plan = &model.Plans{Id: id}
			err = receiver.AddReceiver(Invoker.GetDB())
			if err != nil {
				resp.Code = 1
				resp.Msg = err.Error()
			}
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, receiver)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*PlanHandler) UpdatePlan(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var plan model.Plans
	planId := ctx.Param(":planid")
	id, _ := strconv.ParseInt(planId, 10, 64)
	var resp common.Res
	err := ctx.BindJSON(&plan)
	if err == nil {
		plan.Id = id
		err = plan.UpdatePlan(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, plan)
	} else {
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*PlanHandler) DeletePlan(ctx *gin.Context) {
	logger := gobal.GetLogger()
	planId := ctx.Param(":planid")
	id, _ := strconv.ParseInt(planId, 10, 64)
	var Receiver = &model.Plans{}
	var resp common.Res
	err := Receiver.DeletePlan(Invoker.GetDB(), id)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, planId)
	ctx.JSON(http.StatusOK, resp)
}
