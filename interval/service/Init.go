/*
@Time : 2020/9/3 16:45
@Author : wangyl
@File : Init.go
@Software: GoLand
*/
package service

import (
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/service/HttpService"
	"xing-doraemon/interval/service/TimerService"
)

func Init() error {
	/*
		导入数据连接
	*/
	if err := Invoker.Init(); err != nil {
		return err
	}
	/*
		初始化告警发送器
	*/
	TimerService.InitTimerService()

	/*
		开启http Service
	*/
	if err := HttpService.InitHttpService(); err != nil {
		return err
	}
	return nil
}
