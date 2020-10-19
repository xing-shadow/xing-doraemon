/*
@Time : 2020/7/20 10:08
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package alter_gateway

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/service"
)

var configPath string

func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), "prometheus alerts manager")
	a.HelpFlag.Short('h')
	a.Flag("conf", "config file path").Short('c').StringVar(&configPath)
	if err := SetUp(); err != nil {
		gobal.GetLogger().Panic("set up fail: ", err)
	}
	gobal.GetLogger().Info("set up success")

	if err := service.Init(); err != nil {
		gobal.GetLogger().Panic("service init fail: ", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)
	switch <-quit {

	}
}

func SetUp() error {
	if err := gobal.InitAlterGatewayConfig(configPath); err != nil {
		return err
	}
	return nil
}
