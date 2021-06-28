package Sender

import (
	"go.uber.org/zap"
	"xing-doraemon/global"
)

var sendHandle = map[int]Sender{}

func RegisterSendHandle(method int, handle Sender) {
	if _, ok := sendHandle[method]; !ok {
		sendHandle[method] = handle
	}
}

type Sender interface {
	SendAlert(data ReadyToSend)
	SendRecovery(data ReadyToSend)
}

type ReadyToSend struct {
	RulId  uint
	Url    string
	Alerts []SendAlert
}

type SendAlert struct {
	Id       uint
	Value    float64
	Summary  string
	Hostname string
	Labels   string
}

func SendToMethod(data map[int]*ReadyToSend, isRecovery bool) {
	for method, item := range data {
		handler, ok := sendHandle[method]
		if ok {
			if !isRecovery {
				go handler.SendAlert(*item)
			} else {
				go handler.SendRecovery(*item)
			}
		} else {
			global.Log.Error("not found send handler", zap.Int("method", method))
		}
	}
}
