package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"xing-doraemon/configs"
	"xing-doraemon/global"
	"xing-doraemon/internal/app"
	"xing-doraemon/internal/service"
)

var configPath string

// @title prometheus Alert management center
// @version 2.0
func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), "prometheus alerts dynamic config")
	a.HelpFlag.Short('h')
	a.Flag("conf", "config file path").Short('c').StringVar(&configPath)
	if _, err := a.Parse(os.Args[1:]); err != nil {
		panic("parse cmd line fail: " + err.Error())
	}
	if err := configs.InitConfig(configPath); err != nil {
		panic("config init fail: " + err.Error())
	}
	if err := global.InitGlobal(); err != nil {
		panic("global init fail: " + err.Error())
	}
	if err := service.Init(); err != nil {
		panic("service init fail: " + err.Error())
	}

	quit := make(chan os.Signal, 1)
	var exit = make(chan error, 1)
	app.Init(exit)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		fmt.Println("Program Normal Exit")
	case err := <-exit:
		panic("Program Exit:" + err.Error())
	}
}
