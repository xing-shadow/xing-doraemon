/*
 * @Time : 2020/11/17 16:13
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package CasbinService

import (
	"github.com/jinzhu/gorm"
	"xing-doraemon/pkg/setting"
)

var opt Option

type Option struct {
	*gorm.DB
	Config setting.Casbin
}

func Init(option Option) error {
	opt = option
	casbinAdapter := &CasbinAdapter{}
	err := InitCasbin(opt.Config, casbinAdapter)
	if err != nil {
		return err
	}
	return nil
}
