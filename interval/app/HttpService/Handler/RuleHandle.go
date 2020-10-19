/*
 * @Time : 2020/10/19 17:09
 * @Author : wangyl
 * @File : RuleHandle.go
 * @Software: GoLand
 */
package Handler

import (
	"net/http"
	"strconv"
	"xing-doraemon/interval/model/view"
	"xing-doraemon/interval/service/RuleService"
	"xing-doraemon/pkg/app/Resp"
)

func GetRuleAll(ctx *Resp.Context) {
	var param view.GetRulesReq
	err := ctx.BindJSON(&param.PaginationRequest)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := RuleService.GetAllRule(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

func CreateRule(ctx *Resp.Context) {
	var param view.CreateRuleReq
	err := ctx.BindJSON(&param)
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

func ModifyRule(ctx *Resp.Context) {
	var param view.ModifyRuleReq
	ruleID, err := strconv.Atoi(ctx.Param("ruleid"))
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	if ruleID < 1 {
		ctx.ToResponse(Resp.MsgError, "invalid param", ctx.WithStatus(http.StatusOK))
		return
	}
	param.ID = uint(ruleID)
	err = ctx.BindJSON(&param)
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
