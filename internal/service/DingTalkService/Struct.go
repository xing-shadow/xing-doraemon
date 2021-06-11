package DingTalkService

type DingTalk struct {
	Type string       `json:"msgtype"`
	Text MarkDownText `json:"markdown"`
}

type MarkDownText struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Response struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

type DingTalkInfo struct {
	Title     string      `json:"title"`
	Text      string      `json:"text"`
	AlertList []AlertItem `json:"alert_list"`
}

type AlertItem struct {
	Labels string  `json:"labels"`
	Value  float64 `json:"value"`
}
