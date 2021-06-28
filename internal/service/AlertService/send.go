package AlertService

import (
	"fmt"
	"github.com/prometheus/prometheus/rules"
	"runtime"
	"sync"
	"time"
	"xing-doraemon/global"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/service/AlertService/Sender"
	"xing-doraemon/internal/service/AlertService/Sender/DingTalk"
)

const (
	DingTalkHandle = iota
)

func init() {
	Sender.RegisterSendHandle(DingTalkHandle, DingTalk.DingTalkSender{})
}

var recoverySend = map[int]*Sender.ReadyToSend{}
var recoveryMutex sync.Mutex

func alertInspection() {
	for {
		current := time.Now()
		time.Sleep(time.Duration(60-current.Second()) * time.Second)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 16384)
					buf = buf[:runtime.Stack(buf, false)]
					global.Log.Error(fmt.Sprintf("Panic in alertInspection:%v\n%s", err, buf))
				}
			}()
			var alerts []db.Alert
			tx := opt.DB.Begin()
			tx.Select("id,rule_id,value,count,summary,description,hostname,confirmed_before,fired_at,labels").Where("status=? and confirmed_before < ?", rules.StateFiring, time.Now()).Find(&alerts)
			aggregation := make(map[uint][]db.Alert)
			for i := 0; i < len(alerts); i++ {
				aggregation[alerts[i].RuleId] = append(aggregation[alerts[i].RuleId], alerts[i])
			}
			sendData := filterAlert(aggregation)
			Sender.SendToMethod(sendData, false)
			recoveryMutex.Lock()
			recovery2send := recoverySend
			recoverySend = map[int]*Sender.ReadyToSend{}
			recoveryMutex.Unlock()
			Sender.SendToMethod(recovery2send, true)
		}()
	}
}
