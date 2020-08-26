/*
@Time : 2020/7/20 11:23
@Author : wangyl
@File : SetUserRouter.go
@Software: GoLand
*/
package UserRouter

import (
	"github.com/gin-gonic/gin"
)

func SetUserRouter(group *gin.RouterGroup) {
	group.GET("/", handleGet)
	group.POST("/", handlePost)
	group.PUT("/", handlePut)
	group.DELETE("/:id", handDelete)
}
