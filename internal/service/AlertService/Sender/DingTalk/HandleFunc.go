package DingTalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xing-doraemon/global"
	"xing-doraemon/internal/service/AlertService/Sender"
)

type DingTalkSender struct {
}

func (d DingTalkSender) SendAlert(data Sender.ReadyToSend) {

}

func (d DingTalkSender) SendRecovery(data Sender.ReadyToSend) {

}

func HandleDingTalkInfoFunc(args interface{}) {
	data, ok := args.(DingTalkInfo)
	if !ok {
		return
	}
	title := fmt.Sprintf("Prometheus告警")
	text := fmt.Sprintf("### %s PromQL:%s\n", data.Title, data.Text)
	for _, item := range data.AlertList {
		text += fmt.Sprintf("- %s %f\n", item.Labels, item.Value)
	}
	msg := &DingTalk{
		Type: "markdown",
		Text: MarkDownText{
			Title: title,
			Text:  text,
		},
	}
	msgByte, err := json.Marshal(msg)
	if err != nil {
		global.Log.Error("HandleDingTalkInfoFunc json.Marshal fail:" + err.Error())
		return
	}
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(msgByte))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Log.Warn("send ding talk fail:" + err.Error())
	} else {
		fmt.Println(string(body))
	}
}
