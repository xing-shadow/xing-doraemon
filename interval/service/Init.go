/*
@Time : 2020/9/3 16:45
@Author : wangyl
@File : Init.go
@Software: GoLand
*/
package service

import (
	"xing-doraemon/interval/Invoker"
)

func Init() error {
	/*
		导入数据连接
	*/
	if err := Invoker.Init(); err != nil {
		return err
	}
	return nil
}
