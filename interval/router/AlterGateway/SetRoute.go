/*
@Time : 2020/7/20 10:10
@Author : wangyl
@File : SetRoute.go
@Software: GoLand
*/
package AlterGateway

import (
	"xing-doraemon/cmd/alter-gateway/NetService/LoginRouter"
	"xing-doraemon/cmd/alter-gateway/NetService/LogoutRouter"
	"xing-doraemon/cmd/alter-gateway/NetService/RuleRouter"
	"xing-doraemon/cmd/alter-gateway/NetService/UserRouter"
)

func (pThis *NetService) SetRoute() {
	pThis.router.Use(Cors())
	pThis.router.Use(FilterUser())
	api := pThis.router.Group("/api/v1")
	{
		login := api.Group("/login")
		LoginRouter.SetLoginRouter(login)
	}
	{
		logout := api.Group("/logout")
		LogoutRouter.SetLogoutRouter(logout)
	}
	{
		users := api.Group("/users")
		UserRouter.SetUserRouter(users)
	}
	{
		rules := api.Group("/rules")
		RuleRouter.SetRuleRouter(rules)
	}
	{
		plans := api.Group("/plans")
	}
	{
		receivers := api.Group("/receivers")
	}
	{
		groups := api.Group("/groups")
	}
	{
		proms := api.Group("/proms")
	}
	{
		maintains := api.Group("/maintains")
	}
	{
		manages := api.Group("/manages")
	}
	{
		configs := api.Group("/configs")
	}
}
