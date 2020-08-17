/*
@Time : 2020/7/23 15:53
@Author : wangyl
@File : Middleware.go
@Software: GoLand
*/
package NetService

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var AllowFileTypeList = []string{"html","css","js","png","json"}

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
