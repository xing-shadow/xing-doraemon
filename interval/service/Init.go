/*
@Time : 2020/9/3 16:45
@Author : wangyl
@File : Init.go
@Software: GoLand
*/
package service

import (
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/service/PlanService"
	"xing-doraemon/interval/service/PromService"
	"xing-doraemon/interval/service/RuleService"
)

func Init() error {
	/*
		导入数据连接
	*/
	if err := Invoker.Init(); err != nil {
		return err
	}
	RuleService.Init(RuleService.Option{DB: Invoker.MysqlInvoker})

	PlanService.Init(PlanService.Option{DB: Invoker.MysqlInvoker})

	PromService.Init(PromService.Option{DB: Invoker.MysqlInvoker})
	return nil
}
