/*
 * @Time : 2020/10/19 17:00
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package app

import "xing-doraemon/interval/app/HttpService"

func Init() error {
	if err := HttpService.Init(); err != nil {
		return err
	}
	return nil
}
