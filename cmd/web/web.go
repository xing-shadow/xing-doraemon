/*
@Time : 2020/7/20 9:52
@Author : wangyl
@File : main.go.go
@Software: GoLand
*/
package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	httpPort int
)

func init() {
	WebCmd.Flags().IntVarP(&httpPort, "http-port", "", 80, "http端口")
}

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "web display",
	Run: func(cmd *cobra.Command, args []string) {
		WebRunFunc()
	},
}

func WebRunFunc() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("web/dist/", true)))
	if err := router.Run(fmt.Sprintf(":%d", httpPort)); err != nil {
		fmt.Println(err)
	}

}

var AllowFileTypeList = []string{"html", "css", "js", "png", "json"}

func AllowFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Origin", "*")
		FileTypeMatch := false
		uri := c.Request.URL.Path
		parts := strings.Split(uri, ".")
		if len(parts) > 1 {
			strFileType := parts[len(parts)-1]
			for _, fType := range AllowFileTypeList {
				if strFileType == fType {
					FileTypeMatch = true
				}
			}
			if !FileTypeMatch {
				c.AbortWithStatus(http.StatusForbidden)
			}
		} else {
			c.Next()
		}
	}
}
