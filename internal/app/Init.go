/*
 * @Time : 2020/10/19 17:00
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package app

import (
	"xing-doraemon/configs"
	"xing-doraemon/internal/app/HttpService"
)

func Init(exit chan error) {
	go func() {
		if err := HttpService.Init(configs.Cfg.App); err != nil {
			exit <- err
		}
	}()
}
