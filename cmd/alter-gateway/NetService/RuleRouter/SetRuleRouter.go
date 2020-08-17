/*
@Time : 2020/7/20 11:47
@Author : wangyl
@File : SetRuleRouter.go
@Software: GoLand
*/
package RuleRouter

import (
	"github.com/gin-gonic/gin"
)

func SetRuleRouter(group *gin.RouterGroup)  {
	group.GET("/",)
	group.POST("/",)
	group.PUT("/:ruleid",)
	group.DELETE(":ruleid",)
}
