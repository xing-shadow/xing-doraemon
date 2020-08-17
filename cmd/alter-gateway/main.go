/*
@Time : 2020/7/20 10:08
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package main

import (
	"log"
	"os"
	"strconv"

	"xing-doraemon/cmd/alter-gateway/Datamanager"
	"xing-doraemon/cmd/alter-gateway/Datamanager/MongoDataManger"
	"xing-doraemon/cmd/alter-gateway/NetService"
)

var (
	httpPort int
	Info     = log.New(os.Stdout, "\x1b[36m[INFO]\x1b[0m", log.LstdFlags|log.Lshortfile)
	Error    = log.New(os.Stdout, "\x1b[31m[INFO]\x1b[0m", log.LstdFlags|log.Lshortfile)
)

func main() {
	for index, k := range os.Args {
		if k == "--mongo" {
			MongoDataManger.MongoUrl = os.Args[index+1]
		} else if k == "--http-port" {
			httpPort, _ = strconv.Atoi(os.Args[index+1])
		}
	}
	//1. InitDataManager
	Info.Println("Start init DataManger")
	if err := Datamanager.Init(); err != nil {
		Error.Println("Init DataManager fail: ", err)
		os.Exit(-1)
	}
	Info.Println("Start Http Service")
	var httpService NetService.NetService
	if err := httpService.StartWork(httpPort); err != nil {
		Error.Println("init http Service fail")
		os.Exit(1)
	}
}
