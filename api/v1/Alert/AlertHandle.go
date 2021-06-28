package Alert

import (
	"net/http"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/model/view"
	"xing-doraemon/internal/service/AlertService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 获取告警列表
// @Produce  json
// @Param page query int true "页号"
// @Param page_size query int true "页大小"
// @Success 200 {object} Resp.Response
// @Router /api/v1/alerts [get]
func GetAlerts(ctx *Resp.Context) {
	var param view.GetAlertsReq
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := AlertService.GetAlertList(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 告警确认
// @Produce  json
// @Param body body view.ConfirmAlertsReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/alerts/confirm [post]
func ConfirmAlerts(ctx *Resp.Context) {
	var param view.ConfirmAlertsReq
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	u, _ := ctx.Get("user")
	err = AlertService.ConfirmAlertList(u.(*db.User).Name, param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "success", ctx.WithStatus(http.StatusOK))
	return
}
