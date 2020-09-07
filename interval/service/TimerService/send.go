/*
@Time : 2020/9/7 11:44
@Author : wangyl
@File : send.go
@Software: GoLand
*/
package TimerService

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
	"time"

	"xing-doraemon/gobal"
	"xing-doraemon/pkg/common"
	"xing-doraemon/pkg/notify"
)

/*
 send alert if rule is triggered.
*/
func Sender(SendClass map[string][]common.Ready2Send, now string) {
	for k, v := range SendClass {
		switch k {
		case common.AlertMethodSms:
			go SendAll(k, "mis", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
		case common.AlertMethodLanxin:
			go SendAll(k, "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
			//logs.Alertloger.Info("[%s]%v:", now, v)
		case common.AlertMethodCall:
			go SendAll(k, "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, v, now)
		default:
			route3rdChannel(strings.Split(k, " "), v, now)
		}
	}
}

func route3rdChannel(method []string, content []common.Ready2Send, sendTime string) {
	switch method[0] {
	case common.AlertMethodHook:
		go Send2Hook(content, sendTime, "alert", method[1])
	case common.AlertMethodDingTalk:
		go notify.Send2DingTalk(content, false, sendTime, method[1], method[2])
	}
}

/*
 send recovery message if alert recovered.
*/
func RecoverSender(SendClass map[string]map[[2]int64]*common.Ready2Send, now string) {
	lanxin := []common.Ready2Send{}
	for _, v := range SendClass[common.AlertMethodLanxin] {
		lanxin = append(lanxin, *v)
	}
	go SendRecover(gobal.GetAlterGatewayConfig().Send.LanxinUrl, "StreeAlert", map[string]string{"key": "6E358A78-0A5B-49D2-A12F-6A4EB07A9671"}, lanxin, now)
	//logs.Panic.Info("send[%s]:%v", now, lanxin)
	delete(SendClass, common.AlertMethodLanxin)

	for k := range SendClass {
		var hook []common.Ready2Send
		for _, u := range SendClass[k] {
			hook = append(hook, *u)
		}
		// split by space
		t := strings.Split(k, " ")
		// route to special channel
		switch t[0] {
		case common.AlertMethodHook:
			go Send2Hook(hook, now, "recover", t[1])
		case common.AlertMethodDingTalk:
			go notify.Send2DingTalk(hook, true, now, t[1], t[2])
		default:
			go Send2Hook(hook, now, "recover", k[5:])
		}
	}
}

func SendAll(method string, from string, param map[string]string, content []common.Ready2Send, sendTime string) {
	logger := gobal.GetLogger()
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Errorf("Panic in SendAll:%v\n%s", e, buf)
		}
	}()
	if method == common.AlertMethodSms {
		url := gobal.GetAlterGatewayConfig().Send.SmsUrl
		for _, i := range content {
			msg := []string{"[故障:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
			msg = append(msg, "[时间] "+sendTime)
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	} else if method == common.AlertMethodLanxin {
		url := gobal.GetAlterGatewayConfig().Send.LanxinUrl
		for _, i := range content {
			msg := []string{"[故障:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
			for _, j := range i.Alerts {
				duration, _ := time.ParseDuration(strconv.FormatInt(int64(j.Count), 10) + "m")

				id := strconv.FormatInt(j.Id, 10)
				value := strconv.FormatFloat(j.Value, 'f', 2, 64)
				msg = append(msg, "["+duration.String()+"][ID:"+id+"] "+j.Hostname+" 当前值:"+value)
			}
			msg = append(msg, "[时间] "+sendTime)
			msg = append(msg, "[确认链接] "+gobal.GetAlterGatewayConfig().Send.WebUrl+"/alerts_confirm/"+strconv.FormatInt(i.RuleId, 10)+"?start="+strconv.FormatInt(i.Start, 10))
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	} else {
		url := gobal.GetAlterGatewayConfig().Send.CallUrl
		for _, i := range content {
			msg := []string{"故障" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条 " + i.Alerts[0].Summary + " 详细信息请到蓝信查看"}
			data, _ := json.Marshal(common.Msg{
				Content: strings.Join(msg, "\n"),
				From:    from,
				Title:   "Alerts",
				To:      i.User,
			})
			common.HttpPost(url, param, nil, data)
		}
	}
}

type hookRequest struct {
	Type        string               `json:"type"`
	Time        string               `json:"time"`
	RuleId      int64                `json:"rule_id"`
	To          []string             `json:"to"`
	Alerts      []common.SingleAlert `json:"alerts"`
	ConfirmLink string               `json:"confirm_link,omitempty"`
}

func Send2Hook(content []common.Ready2Send, sendTime string, t string, url string) {
	logger := gobal.GetLogger()
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Errorf("Panic in Send2Hook:%v\n%s", e, buf)
		}
	}()
	if t == "recover" {
		for _, i := range content {
			data, _ := json.Marshal(hookRequest{
				Type:   t,
				RuleId: i.RuleId,
				Time:   sendTime,
				To:     i.User,
				Alerts: i.Alerts,
			})
			common.HttpPost(url, nil, common.GenerateJsonHeader(), data)
		}
	} else {
		for _, i := range content {
			data, _ := json.Marshal(hookRequest{
				Type:        t,
				RuleId:      i.RuleId,
				Time:        sendTime,
				ConfirmLink: gobal.GetAlterGatewayConfig().Send.WebUrl + "/alerts_confirm/" + strconv.FormatInt(i.RuleId, 10) + "?start=" + strconv.FormatInt(i.Start, 10),
				To:          i.User,
				Alerts:      i.Alerts,
			})
			common.HttpPost(url, nil, common.GenerateJsonHeader(), data)
		}
	}
}

func SendRecover(url string, from string, param map[string]string, content []common.Ready2Send, sendTime string) {
	logger := gobal.GetLogger()
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Errorf("Panic in SendRecover:%v\n%s", e, buf)
		}
	}()
	for _, i := range content {
		msg := []string{"[故障恢复:" + strconv.FormatInt(int64(len(i.Alerts)), 10) + "条] " + i.Alerts[0].Summary}
		for _, j := range i.Alerts {
			duration, _ := time.ParseDuration(strconv.FormatInt(int64(j.Count), 10) + "m")

			id := strconv.FormatInt(j.Id, 10)
			value := strconv.FormatFloat(j.Value, 'f', 2, 64)
			msg = append(msg, "["+duration.String()+"][ID:"+id+"] "+j.Hostname+" 当前值:"+value)
		}
		msg = append(msg, "[时间] "+sendTime)
		data, _ := json.Marshal(common.Msg{
			Content: strings.Join(msg, "\n"),
			From:    from,
			Title:   "Alerts",
			To:      i.User})
		common.HttpPost(url, param, nil, data)
	}
}
