/*
@Time : 2020/7/20 10:08
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package alter_gateway

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/service"
)

func init() {
	AlterGatewayCmd.Flags().StringVarP(&configPath, "conf-path", "c", "", "配置文件路径")
}

var configPath string
var AlterGatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "alter gateway service",
	Run: func(cmd *cobra.Command, args []string) {
		AlterGatewayRunFunc()
	},
}

func AlterGatewayRunFunc() {
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
	if err := gobal.InitLdap(gobal.GetAlterGatewayConfig().Ldap); err != nil {
		return err
	}
	return nil
}
