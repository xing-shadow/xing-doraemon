/*
@Time : 2020/7/20 10:08
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"xing-doraemon/interval/app"

	"xing-doraemon/global"
	"xing-doraemon/interval/service"
)

var configPath string

func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), "prometheus alerts manager")
	a.HelpFlag.Short('h')
	a.Flag("conf", "config file path").Short('c').StringVar(&configPath)
	if err := SetUp(); err != nil {
		global.GetLogger().Panic("set up fail: ", err)
	}
	global.GetLogger().Info("set up success")

	if err := service.Init(); err != nil {
		global.GetLogger().Panic("service init fail: ", err)
	}
	if err := app.Init(); err != nil {
		global.GetLogger().Panic("app init fail: ", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)
	switch <-quit {

	}
}

func SetUp() error {
	if err := global.InitAlterGatewayConfig(configPath); err != nil {
		return err
	}
	return nil
}
