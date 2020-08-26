/*
@Time : 2020/8/24 11:49
@Author : wangyl
@File : root.go
@Software: GoLand
*/
package cmd

import (
	"github.com/spf13/cobra"
	alter_gateway "xing-doraemon/cmd/alter-gateway"
	"xing-doraemon/cmd/web"
)

func init() {
	rootCmd.AddCommand(web.WebCmd)
	rootCmd.AddCommand(alter_gateway.AlterGatewayCmd)
}

var rootCmd = &cobra.Command{}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	} else {
		return nil
	}
}
