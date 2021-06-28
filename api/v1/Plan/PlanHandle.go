package Plan

import (
	"net/http"
	"xing-doraemon/internal/model/view"
	"xing-doraemon/internal/service/PlanService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 获取单个Plan
// @Produce  json
// @Param id query string true "序号"
// @Success 200 {object} Resp.Response
// @Router /api/v1/plan [get]
func GetPlan(ctx *Resp.Context) {
	var param view.GetPlan
	err := ctx.BindQuery(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := PlanService.GetPlan(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 获取所有prom名
// @Produce  json
// @Success 200 {object} Resp.Response
// @Router /api/v1/plan/allName [get]
func GetPlanAllName(ctx *Resp.Context) {
	data, err := PlanService.GetPlanAllName()
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 获取Plan列表，分页
// @Produce  json
// @Param page query string true "页序号"
// @Param page_size query string true "页大小"
// @Success 200 {object} Resp.Response
// @Router /api/v1/plans [get]
func GetPlanPagination(ctx *Resp.Context) {
	var param view.GetPlanList
	err := ctx.BindQuery(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := PlanService.GetPlanPagination(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 创建plan
// @Produce  json
// @Param body body view.CreatePlanReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/plan/add [post]
func CreatePlan(ctx *Resp.Context) {
	var param view.CreatePlanReq
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = PlanService.CreatePlan(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 修改plan
// @Produce  json
// @Param body body view.ModifyPlanReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/plan/update [post]
func ModifyPlan(ctx *Resp.Context) {
	var param view.ModifyPlanReq
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = PlanService.ModifyPlan(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 删除plan
// @Produce  json
// @Param body body view.DeleteRuleReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/plan/delete [post]
func DeletePlan(ctx *Resp.Context) {
	var param view.DeleteRuleReq
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = PlanService.DeletePlan(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}
