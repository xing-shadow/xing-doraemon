/*
@Time : 2020/7/23 16:35
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"iPublic/LoggerModular"

	"xing-doraemon/cmd/Hook/NetService"
)

var HttpPort int

func main() {
	LoggerModular.SetLogLevel(logrus.InfoLevel)
	for index, k := range os.Args {
		if k == "--http-port" {
			HttpPort,_ = strconv.Atoi(os.Args[index+1])
		}
	}
	var httpService NetService.NetService
	if err := httpService.StartWork(HttpPort);err != nil {
		fmt.Println(err)
	}
}



