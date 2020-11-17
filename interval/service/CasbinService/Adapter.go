/*
 * @Time : 2020/11/17 16:17
 * @Author : wangyl
 * @File : Adapter.go
 * @Software: GoLand
 */
package CasbinService

import (
	"fmt"
	CasbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"xing-doraemon/global"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

type CasbinAdapter struct{}

func (c *CasbinAdapter) LoadPolicy(model CasbinModel.Model) (err error) {
	err = c.loadPolicyAuth(model)
	if err != nil {
		global.GetLogger().Error("load policy auth error:", err)
		return
	}
	err = c.loadPolicyGroup(model)
	if err != nil {
		global.GetLogger().Error("load policy group error:", err)
		return err
	}
	return nil
}

func (c *CasbinAdapter) loadPolicyAuth(model CasbinModel.Model) (err error) {
	authList, err := PolicyAuthList()
	if err != nil {
		return
	}
	for _, item := range authList {
		line := fmt.Sprintf("p,%s,%s,%s", item.Sub, item.Obj, item.Act)
		persist.LoadPolicyLine(line, model)
	}
	return
}

func (c *CasbinAdapter) loadPolicyGroup(model CasbinModel.Model) (err error) {
	authList, err := PolicyGroupList()
	if err != nil {
		return
	}
	var line string
	for _, item := range authList {
		line = fmt.Sprintf("g,%s,%s", item.UserName, item.GroupName)
		persist.LoadPolicyLine(line, model)
	}
	return
}

func (c *CasbinAdapter) SavePolicy(model CasbinModel.Model) error {
	return nil
}

func (c *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

func (c *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

func (c *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
