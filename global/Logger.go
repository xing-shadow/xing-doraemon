/*
@Time : 2020/8/24 18:17
@Author : wangyl
@File : Logger.go
@Software: GoLand
*/
package global

import (
	"github.com/sirupsen/logrus"

	"xing-doraemon/pkg/Logger"
)

var logger = Logger.NewLogger("./", logrus.InfoLevel, false)

func GetLogger() *logrus.Logger {
	return logger.Log()
}
