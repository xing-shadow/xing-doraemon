package Prom

import (
	"net/http"
	"xing-doraemon/internal/model/view"
	"xing-doraemon/internal/service/PromService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 获取proms, 分页
// @Produce  json
// @Param page query string true "页序号"
// @Param page_size query string true "页大小"
// @Success 200 {object} Resp.Response
// @Router /api/v1/proms [get]
func GetPromsPagination(ctx *Resp.Context) {
	var param view.GetProms
	err := ctx.BindQuery(&param)
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
// @Router /api/v1/prom [get]
func GetProm(ctx *Resp.Context) {
	var param view.GetProm
	err := ctx.BindQuery(&param)
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

// @Summary 获取所有prom名
// @Produce  json
// @Success 200 {object} Resp.Response
// @Router /api/v1/prom/allName [get]
func GetPromAllName(ctx *Resp.Context) {
	data, err := PromService.GetPromAllName()
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
// @Router /api/v1/prom/add [post]
func CreateProm(ctx *Resp.Context) {
	var param view.CreateProm
	err := ctx.BindJSON(&param)
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
// @Router /api/v1/prom/update [post]
func ModifyProm(ctx *Resp.Context) {
	var param view.ModifyProm
	err := ctx.BindJSON(&param)
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
// @Router /api/v1/prom/delete [post]
func DeleteProm(ctx *Resp.Context) {
	var param view.DeleteProm
	err := ctx.BindJSON(&param)
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
