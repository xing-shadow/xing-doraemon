/*
@Time : 2020/7/23 16:37
@Author : wangyl
@File : Struce.go
@Software: GoLand
*/
package NetService

type Hook struct {
	Type        string   `json:"type"`
	Time        string   `json:"time"`
	RuleID      int      `json:"rule_id"`
	To          []string `json:"to"`
	ConfirmLink string   `json:"confirm_link"`
	Alerts      []Alerts `json:"alerts"`
}

type Alerts struct {
	ID          int                    `json:"id"`
	Count       int                    `json:"count"`
	Value       float64                `json:"value"`
	Summary     string                 `json:"summary"`
	Description string                 `json:"description"`
	Hostname    string                 `json:"hostname"`
	Labels      map[string]interface{} `json:"labels"`
}
