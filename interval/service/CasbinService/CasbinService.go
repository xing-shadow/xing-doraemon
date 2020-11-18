/*
 * @Time : 2020/11/17 16:15
 * @Author : wangyl
 * @File : CasbinService.go
 * @Software: GoLand
 */
package CasbinService

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"xing-doraemon/global"
	"xing-doraemon/pkg/setting"
	"xing-doraemon/pkg/xtime"
)

type CasbinService struct {
	*casbin.SyncedEnforcer
	enabled bool
}

var (
	Casbin *CasbinService
)

// InitCasbin 初始化casbin

func InitCasbin(conf setting.Casbin, adapter persist.Adapter) (err error) {
	Casbin = &CasbinService{
		SyncedEnforcer: nil,
		enabled:        true,
	}
	if !conf.Enable {
		return
	}
	if conf.Model == "" {
		Casbin.SyncedEnforcer = new(casbin.SyncedEnforcer)
	}
	e, err := casbin.NewSyncedEnforcer(conf.Model)
	if err != nil {
		return
	}
	e.EnableLog(conf.Debug)
	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	if err != nil {
		return
	}
	e.EnableEnforce(conf.Enable)
	if conf.AutoLoad {
		e.StartAutoLoadPolicy(xtime.ToDuration(conf.AutoLoadInterval))
	}
	Casbin.SyncedEnforcer = e
	return
}

func (c *CasbinService) CheckPermission(sub, object, action string) (ok bool, err error) {
	if !c.enabled {
		return true, nil
	}
	ok, err = c.Enforce(sub, object, action)
	if err != nil {
		global.GetLogger().Errorf("CasbinService.CheckUserPermission: %s", err.Error())
		return false, err
	}
	return
}
