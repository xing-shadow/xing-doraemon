package service

import (
	"xing-doraemon/configs"
	"xing-doraemon/internal/Invoker"
	"xing-doraemon/internal/service/AlertService"
	"xing-doraemon/internal/service/DingTalkService"
	"xing-doraemon/internal/service/PlanService"
	"xing-doraemon/internal/service/PromService"
	"xing-doraemon/internal/service/RuleEngine"
	"xing-doraemon/internal/service/RuleService"
	"xing-doraemon/internal/service/UserService"
	"xing-doraemon/pkg/xtime"
)

func Init() error {
	/*
		导入数据连接
	*/
	if err := Invoker.Init(configs.Cfg.Mysql); err != nil {
		return err
	}
	UserService.Init(UserService.Option{DB: Invoker.MysqlInvoker, SessCfg: configs.Cfg.Session})

	RuleService.Init(RuleService.Option{DB: Invoker.MysqlInvoker})

	PlanService.Init(PlanService.Option{DB: Invoker.MysqlInvoker})

	PromService.Init(PromService.Option{DB: Invoker.MysqlInvoker})

	if err := DingTalkService.Init(DingTalkService.Option{
		PushAddr: configs.Cfg.Send.WebHook,
	}); err != nil {
		return err
	}

	if err := AlertService.Init(AlertService.Option{DB: Invoker.MysqlInvoker}); err != nil {
		return err
	}

	RuleEngine.Init(RuleEngine.Option{
		DB: Invoker.MysqlInvoker,
		Cfg: RuleEngine.Config{
			NotifyRetries:      configs.Cfg.RuleEngine.NotifyRetries,
			EvaluationInterval: xtime.ToDuration(configs.Cfg.RuleEngine.EvaluationInterval),
			ReloadInterval:     xtime.ToDuration(configs.Cfg.RuleEngine.ReloadInterval),
		},
	})
	return nil
}
