/*
@Time : 2020/9/3 16:45
@Author : wangyl
@File : Init.go
@Software: GoLand
*/
package service

import (
	"xing-doraemon/global"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/service/AlertService"
	"xing-doraemon/interval/service/DingTalkService"
	"xing-doraemon/interval/service/PlanService"
	"xing-doraemon/interval/service/PromService"
	"xing-doraemon/interval/service/RuleEngine"
	"xing-doraemon/interval/service/RuleService"
	"xing-doraemon/interval/service/UserService"
	"xing-doraemon/pkg/xtime"
)

func Init() error {
	/*
		导入数据连接
	*/
	if err := Invoker.Init(); err != nil {
		return err
	}
	UserService.Init(UserService.Option{DB: Invoker.MysqlInvoker})

	RuleService.Init(RuleService.Option{DB: Invoker.MysqlInvoker})

	PlanService.Init(PlanService.Option{DB: Invoker.MysqlInvoker})

	PromService.Init(PromService.Option{DB: Invoker.MysqlInvoker})
	if err := DingTalkService.Init(DingTalkService.Option{
		PushAddr: global.GetAlterGatewayConfig().Send.WebHook,
	}); err != nil {
		return err
	}
	if err := AlertService.Init(AlertService.Option{DB: Invoker.MysqlInvoker}); err != nil {
		return err
	}

	if global.GetAlterGatewayConfig().RuleEngine.Enable {
		RuleEngine.Init(RuleEngine.Option{
			DB: Invoker.MysqlInvoker,
			Cfg: RuleEngine.Config{
				NotifyRetries:      global.GetAlterGatewayConfig().RuleEngine.NotifyRetries,
				EvaluationInterval: xtime.ToDuration(global.GetAlterGatewayConfig().RuleEngine.EvaluationInterval),
				ReloadInterval:     xtime.ToDuration(global.GetAlterGatewayConfig().RuleEngine.ReloadInterval),
			},
		})
	}

	return nil
}
