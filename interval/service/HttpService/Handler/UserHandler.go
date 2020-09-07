/*
@Time : 2020/9/7 15:29
@Author : wangyl
@File : UserHandler.go
@Software: GoLand
*/
package Handler

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type UserHandler struct {
}

func (u *UserHandler) GetAllUsers(ctx *gin.Context) {
	var Receiver = new(model.Users)
	users := Receiver.GetAll(Invoker.GetDB())
	resp := &common.Res{
		Code: 0,
		Msg:  "",
		Data: users,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserHandler) AddUser(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var userInfo model.Users
	var resp common.Res
	if err := ctx.BindJSON(&userInfo); err != nil {
		logger.Error("Unmarshal plan error: ", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		err = userInfo.AddUser(Invoker.GetDB())
		if err != nil {
			resp.Code = 1
			resp.Msg = err.Error()
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, userInfo)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserHandler) UpdatePassword(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var newInfo struct {
		Name               string `json:"name"`
		OldPassword        string `json:"oldpassword"`
		NewPassword        string `json:"newpassword"`
		ConfirmNewPassword string `json:"confirmnewpassword"`
	}
	var resp common.Res
	var userInfo model.Users
	if err := ctx.BindJSON(&newInfo); err != nil {
		logger.Error("Unmarshal plan error: ", err)
		resp.Code = 1
		resp.Msg = "Unmarshal error"
	} else {
		if newInfo.ConfirmNewPassword == newInfo.NewPassword {
			if sessions.Default(ctx).Get("username") == newInfo.Name {
				err = userInfo.UpdatePassword(Invoker.GetDB(), newInfo.Name, newInfo.OldPassword, newInfo.NewPassword)
				if err != nil {
					resp.Code = 1
					resp.Msg = err.Error()
				}
			} else {
				resp.Code = 1
				resp.Msg = "Inconsistent user identity"
			}
		} else {
			resp.Code = 1
			resp.Msg = "The two new passwords are inconsistent"
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, newInfo)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (u *UserHandler) DeleteUsers(ctx *gin.Context) {
	logger := gobal.GetLogger()
	userId := ctx.Param("id")
	var userInfo model.Users
	var resp common.Res
	err := userInfo.DeleteUsers(Invoker.GetDB(), userId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Info("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, userId)
	ctx.JSON(http.StatusOK, resp)
}
