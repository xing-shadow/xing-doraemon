/*
@Time : 2020/7/20 10:08
@Author : wangyl
@File : main.go
@Software: GoLand
*/
package alter_gateway

import (
	"github.com/spf13/cobra"

	"xing-doraemon/gobal"
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
