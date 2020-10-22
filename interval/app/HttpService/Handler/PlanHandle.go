/*
 * @Time : 2020/10/22 14:49
 * @Author : wangyl
 * @File : PlanHandle.go
 * @Software: GoLand
 */
package Handler

import (
	"net/http"
	"xing-doraemon/interval/model/view"
	"xing-doraemon/interval/service/PlanService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 获取单个Plan
// @Produce  json
// @Param id query string true "序号"
// @Success 200 {object} Resp.Response
// @Router /api/v1/planID [get]
func GetPlan(ctx *Resp.Context) {
	var param view.GetPlan
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	if param.Id <= 0 {
		ctx.ToResponse(Resp.MsgError, "invalid params", ctx.WithStatus(http.StatusOK))
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

// @Summary 获取Plan列表，分页
// @Produce  json
// @Param page query string true "页序号"
// @Param page_size query string true "页大小"
// @Success 200 {object} Resp.Response
// @Router /api/v1/plan [get]
func GetPlanPagination(ctx *Resp.Context) {
	var param view.GetPlanList
	err := ctx.BindParam(&param)
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

// @Summary 获取所有Plan
// @Produce  json
// @Success 200 {object} Resp.Response
// @Router /api/v1/plans [get]
func GetAllPlan(ctx *Resp.Context) {
	data, err := PlanService.GetAllPlan()
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
// @Router /api/v1/plan [post]
func CreatePlan(ctx *Resp.Context) {
	var param view.CreatePlanReq
	err := ctx.BindParam(&param)
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
// @Router /api/v1/plan [put]
func ModifyPlan(ctx *Resp.Context) {
	var param view.ModifyPlanReq
	err := ctx.BindParam(&param)
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
// @Router /api/v1/plan [delete]
func DeletePlan(ctx *Resp.Context) {
	var param view.DeleteRuleReq
	err := ctx.BindParam(&param)
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
