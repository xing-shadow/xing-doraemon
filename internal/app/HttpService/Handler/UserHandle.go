/*
 * @Time : 2020/11/13 17:32
 * @Author : wangyl
 * @File : UserHandle.go
 * @Software: GoLand
 */
package Handler

import (
	"net/http"
	"xing-doraemon/internal/model/view"
	"xing-doraemon/internal/service/UserService"
	"xing-doraemon/pkg/App/Resp"
)

// @Summary 用户登录
// @Produce  json
// @Param body body view.LoginReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/rule [post]
func UserLogin(ctx *Resp.Context) {
	var param view.LoginReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := UserService.Login(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 用户列表
// @Produce  json
// @Param page query string true "序号"
// @Param page_size query string true "序号"
// @Success 200 {object} Resp.Response
// @Router /api/v1/user/list [get]
func UserList(ctx *Resp.Context) {
	var param view.UserListReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	data, err := UserService.UserList(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK), ctx.WithData(data))
	return
}

// @Summary 添加用户
// @Produce  json
// @Param body body view.UserCreateReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/user/create [post]
func UserCreate(ctx *Resp.Context) {
	var param view.UserCreateReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = UserService.CreateUser(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 修改用户
// @Produce  json
// @Param body body view.UserUpdateReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/user/update [put]
func UserUpdate(ctx *Resp.Context) {
	var param view.UserUpdateReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = UserService.UpdateUser(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK))
	return
}

// @Summary 修改用户
// @Produce  json
// @Param body body view.UserDeleteReq true "body"
// @Success 200 {object} Resp.Response
// @Router /api/v1/user/delete [delete]
func UserDelete(ctx *Resp.Context) {
	var param view.UserDeleteReq
	err := ctx.BindParam(&param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	err = UserService.DeleteUser(param)
	if err != nil {
		ctx.ToResponse(Resp.MsgError, err.Error(), ctx.WithStatus(http.StatusOK))
		return
	}
	ctx.ToResponse(Resp.MsgOk, "Success", ctx.WithStatus(http.StatusOK))
	return
}
