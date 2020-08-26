/*
@Time : 2020/7/20 10:08
@Author : wangyl
@File : NetService.go
@Software: GoLand
*/
package AlterGateway

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NetService struct {
	router *gin.Engine
}

func (pThis *NetService) Init() error {

	return errors.New("NetService Init Failed! ")
}

func (pThis *NetService) StartWork(nPort int) error {
	gin.SetMode(gin.ReleaseMode)
	pThis.router = gin.Default()
	pThis.router.NoMethod(func(ctx *gin.Context) {
		ctx.Data(http.StatusMethodNotAllowed, "text/plain", []byte("Method Not Allowed"))
		ctx.Abort()
	})
	pThis.router.NoRoute(func(ctx *gin.Context) {
		ctx.Data(http.StatusNotFound, "text/plain", []byte("404 page not found"))
		ctx.Abort()
	})

	/*
		设置路由
	*/
	pThis.SetRoute()
	return pThis.router.Run(fmt.Sprintf(":%d", nPort))
}

func (pThis *NetService) StopWork() error {

	return errors.New("NetService Init Failed! ")
}

func (pThis *NetService) UnInit() error {
	return errors.New("NetService UnInit Failed! ")
}

func (pThis *NetService) GetObjectType() (iObjectType, iObjectVersion int) {
	return 1000, 1
}
