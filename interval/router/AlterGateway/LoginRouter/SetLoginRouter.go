/*
@Time : 2020/7/20 11:15
@Author : wangyl
@File : Login.go
@Software: GoLand
*/
package LoginRouter

import (
	"github.com/gin-gonic/gin"
)

func SetLoginRouter(group *gin.RouterGroup) {
	group.GET("/method", MethodHandle)
	group.GET("/username", UsernameHandle)
	group.POST("/local", localHandle)
}
