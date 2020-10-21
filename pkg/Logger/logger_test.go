/*
@Time : 2020/8/24 16:08
@Author : wangyl
@File : logger_test.go.go
@Software: GoLand
*/
package Logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestLogger_Log(t *testing.T) {
	logger := NewLogger("./", logrus.InfoLevel, false)
	logger.Log().WithField("test", "hello").Info("haha")

	logger.Log().Fatal("fadasd")
	for {
		fmt.Println(1)
		time.Sleep(time.Second * 10)
	}
}
