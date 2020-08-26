/*
@Time : 2020/7/20 11:21
@Author : wangyl
@File : SetLogoutRouter.go
@Software: GoLand
*/
package LogoutRouter

import (
	"github.com/gin-gonic/gin"
)

func SetLogoutRouter(groups *gin.RouterGroup) {
	groups.GET("/", handle)
}
