/*
@Time : 2020/8/24 15:34
@Author : wangyl
@File : formart.go
@Software: GoLand
*/
package Logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	//defaultTimestampFormat = time.RFC3339
	defaultTimestampFormat = "2006-01-02T15:04:05"
	red                    = 31
	green                  = 32
	yellow                 = 33
	cyan                   = 34
	magenta                = 35
	blue                   = 36
	gray                   = 37
)

type CustomFormat struct {
	timeStampFormat string
	discolor        bool
}

func (f *CustomFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	var levelText string
	var colorCode int
	var caller string
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	switch entry.Level {
	case logrus.DebugLevel:
		colorCode = gray
	case logrus.InfoLevel:
		colorCode = blue
	case logrus.WarnLevel:
		colorCode = green
	case logrus.ErrorLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.PanicLevel:
		fallthrough
	default:
		colorCode = red
	}
	levelText = strings.ToUpper(entry.Level.String())
	if entry.HasCaller() {
		parts := strings.Split(entry.Caller.File, "/")
		fileName := parts[len(parts)-1]
		funcName := entry.Caller.Function[strings.Index(entry.Caller.Function, ".")+1:]
		line := entry.Caller.Line
		caller = fmt.Sprintf("%s:%d==>%s", fileName, line, funcName)
	}
	if f.discolor {
		fmt.Fprint(b, fmt.Sprintf("[%s] %s ",
			levelText, entry.Message))
	} else {
		fmt.Fprint(b, fmt.Sprintf("\x1b[%dm[%s]\u001B[0m %s ",
			colorCode, levelText, entry.Message))
	}
	if len(entry.Data) != 0 {
		fmt.Fprint(b, "< ")
		for k, v := range entry.Data {
			fmt.Fprint(b, k, "=", v, " ")
		}
		fmt.Fprint(b, ">")
	}
	b.WriteByte('\n')
	fmt.Fprint(b, fmt.Sprintf("	%s %s",
		time.Now().Format(f.timeStampFormat), caller))
	b.WriteByte('\n')
	return b.Bytes(), nil
}
