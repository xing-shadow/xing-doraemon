package Hook

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	pThis.router.POST("/doraemon/hook", HookHandle)
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
