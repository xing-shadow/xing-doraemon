/*
@Time : 2020/8/24 15:30
@Author : wangyl
@File : logger.go
@Software: GoLand
*/
package Logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger  *logrus.Logger
	OldData string
	Path    string
	Name    string
	outFile *os.File
}

func NewLogger(path string, level logrus.Level, discolor bool) *Logger {
	logger := &Logger{}
	logger.Path = path
	logger.Name = filepath.Base(os.Args[0])
	logger.logger = &logrus.Logger{
		Formatter: &CustomFormat{
			timeStampFormat: defaultTimestampFormat,
			discolor:        discolor,
		},
		Out:          os.Stderr,
		ReportCaller: true,
		Level:        level,
	}
	return logger
}

func (l *Logger) Log() *logrus.Logger {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic in Log: ", err)
		}
	}()
	if runtime.GOOS != "linux" {
		return l.logger
	} else {
		year, month, day := time.Now().Date()
		Data := fmt.Sprintf("%s-%d-%2d-%2d", l.Name, year, month, day)
		if Data != l.OldData {
			l.OldData = Data
			fileDir := filepath.Join(l.Path, "logs")
			if _, err := os.Stat(fileDir); os.IsNotExist(err) {
				if errDir := os.Mkdir(fileDir, 0755); errDir != nil {
					fmt.Println("Create logs file Failed, will start without log file: ", errDir)
					return l.logger
				}
			}
			fileName := filepath.Join(fileDir, Data) + ".log"
			if file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
				fmt.Println("Start log file err: ", err)
				return l.logger
			} else {
				if l.outFile != nil {
					l.outFile.Close()
				}
				l.outFile = file
				l.logger.SetOutput(file)
			}
		}
		return l.logger
	}
}
