/*
@Time : 2020/7/20 11:39
@Author : wangyl
@File : Handle.go
@Software: GoLand
*/
package UserRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"xing-doraemon/cmd/alter-gateway/Common"
	"xing-doraemon/cmd/alter-gateway/model"
)

func handleGet(ctx *gin.Context) {
	//TODO 获取所有用户
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "OK"
	resp.Data = []model.Users{
		{
			Id:   0,
			Name: "xing",
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

func handlePost(ctx *gin.Context) {
	//TODO 添加用户
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "OK"
	ctx.JSON(http.StatusOK,resp)
}

func handlePut(ctx *gin.Context)  {
	//TODO 修改用户
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "OK"
	ctx.JSON(http.StatusOK,resp)
}

func handDelete(ctx *gin.Context)  {
	//TODO 删除用户
	var resp Common.Res
	resp.Code = 0
	resp.Msg = "OK"
	ctx.JSON(http.StatusOK,resp)
}