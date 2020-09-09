/*
@Time : 2020/9/7 19:46
@Author : wangyl
@File : ReceiversHandler.go
@Software: GoLand
*/
package Handler

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type ReceiversHandler struct {
}

func (*ReceiversHandler) UpdateReceiver(ctx *gin.Context) {
	logger := gobal.GetLogger()
	var receiver = &model.Receivers{}
	var resp common.Res
	receiverId := ctx.Param("receiverid")
	err := ctx.BindJSON(&receiver)
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
				id, _ := strconv.ParseInt(receiverId, 10, 64)
				receiver.Id = id
				err = receiver.UpdateReceiver(Invoker.GetDB())
				if err != nil {
					resp.Code = 1
					resp.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(receiverId, 10, 64)
			receiver.Id = id
			err = receiver.UpdateReceiver(Invoker.GetDB())
			if err != nil {
				resp.Code = 1
				resp.Msg = err.Error()
			}
		}
		logger.Infof("%s %s %s %v", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, receiver)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (*ReceiversHandler) DeleteReceiver(ctx *gin.Context) {
	logger := gobal.GetLogger()
	receiverId := ctx.Param("receiverid")
	var Receiver = &model.Receivers{}
	var resp common.Res
	err := Receiver.DeleteReceiver(Invoker.GetDB(), receiverId)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
	}
	logger.Infof("%s %s %s %s", sessions.Default(ctx).Get("username"), ctx.Request.RequestURI, ctx.Request.Method, receiverId)
	ctx.JSON(http.StatusOK, resp)
}
