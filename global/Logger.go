package global

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"xing-doraemon/configs"

	"xing-doraemon/pkg/Logger"
)

var Log *zap.Logger

func InitGlobal() (err error) {
	err = Logger.Init(
		Logger.SetFileName(configs.Cfg.Logger.FiLeName),
		Logger.SetLogFileDir(configs.Cfg.Logger.LogDir),
		Logger.SetMaxAge(configs.Cfg.Logger.MaxAge),
		Logger.SetLevel(configs.Cfg.Logger.Level),
		Logger.SetDevelopment(configs.Cfg.Logger.Development),
		Logger.SetMaxSize(configs.Cfg.Logger.MaxSize),
		Logger.SetMaxBackups(configs.Cfg.Logger.MaxBackups),
	)
	if err != nil {
		return errors.Wrap(err, "Init logger")
	}
	Log = Logger.GetLogger()
	return nil
}
