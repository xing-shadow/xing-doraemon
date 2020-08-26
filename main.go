/*
@Time : 2020/8/24 11:45
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package main

import (
	"xing-doraemon/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic("cmd execute fail: " + err.Error())
	}
}
