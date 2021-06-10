/*
 * @Time : 2020/10/19 17:09
 * @Author : wangyl
 * @File : RuleHandle.go
 * @Software: GoLand
 */
package Rule

import (
	"net/http"
	"xing-doraemon/internal/model/view"
	"xing-doraemon/internal/service/RuleService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 获取单个rule
// @Produce  json
// @Param id query string true "序号"
// @Success 200 {object} Resp.Response
// @Router /api/v1/ruleId [get]
func GetRule(ctx *Resp.Context) {
	var param view.GetRule
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	if param.Id <= 0 {
		ctx.ToResponse(Resp.MsgError, "invalid params", ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := RuleService.GetRule(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 获取rules列表，分页
// @Produce  json
// @Param page query string true "序号"
// @Param page_size query string true "序号"
// @Success 200 {object} Resp.Response
// @Router /api/v1/rule [get]
func GetRulePagination(ctx *Resp.Context) {
	var param view.GetRulesReq
	err := ctx.BindParam(&param.PaginationRequest)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := RuleService.GetPaginationRule(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 创建rule
// @Produce  json
// @Param body body view.CreateRuleReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/rule [post]
func CreateRule(ctx *Resp.Context) {
	var param view.CreateRuleReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = RuleService.CreateRule(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 修改rule
// @Produce  json
// @Param body body view.ModifyRuleReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/rule [put]
func ModifyRule(ctx *Resp.Context) {
	var param view.ModifyRuleReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}

	err = RuleService.ModifyRule(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 删除rule
// @Produce  json
// @Param body body view.DeleteRuleReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/rule [delete]
func DeleteRule(ctx *Resp.Context) {
	var param view.DeleteRuleReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = RuleService.DeleteRule(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK))
	return
}
