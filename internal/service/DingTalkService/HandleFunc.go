/*
 * @Time : 2020/11/12 16:45
 * @Author : wangyl
 * @File : HandleFunc.go
 * @Software: GoLand
 */
package DingTalkService

import (
	"encoding/json"
	"fmt"
	"xing-doraemon/global"
)

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
		global.GetLogger().Errorf("HandleDingTalkInfoFunc json.Marshal fail:%v", err)
		return
	}
	resp, err := opt.API.R().SetHeader("Content-Type", "application/json").SetBody(msgByte).Post(opt.PushAddr)
	if err != nil {
		global.GetLogger().Warn("send ding talk fail:", err)
	} else {
		fmt.Println(resp)
	}
}
