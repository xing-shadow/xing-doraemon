/*
 * @Time : 2020/10/22 16:18
 * @Author : wangyl
 * @File : PromHandle.go
 * @Software: GoLand
 */
package Handler

import (
	"net/http"
	"xing-doraemon/interval/model/view"
	"xing-doraemon/interval/service/PromService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 获取proms, 分页
// @Produce  json
// @Param page query string true "页序号"
// @Param page_size query string true "页大小"
// @Success 200 {object} Resp.Response
// @Router /api/v1/prom [get]
func GetPromsPagination(ctx *Resp.Context) {
	var param view.GetProms
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := PromService.GetPromPagination(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))

	return
}

// @Summary 获取单个prom
// @Produce  json
// @Param Id query string true "页序号"
// @Success 200 {object} Resp.Response
// @Router /api/v1/promId [get]
func GetProm(ctx *Resp.Context) {
	var param view.GetProm
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := PromService.GetProm(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))

	return
}

// @Summary 创建prom
// @Produce  json
// @Param body body view.CreateProm true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/prom [post]
func CreateProm(ctx *Resp.Context) {
	var param view.CreateProm
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = PromService.CreateProms(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 修改prom
// @Produce  json
// @Param body body view.ModifyProm true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/prom [put]
func ModifyProm(ctx *Resp.Context) {
	var param view.ModifyProm
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = PromService.ModifyProm(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 删除prom
// @Produce  json
// @Param body body view.DeleteProm true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/prom [delete]
func DeleteProm(ctx *Resp.Context) {
	var param view.DeleteProm
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = PromService.DeleteProm(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}
