/*
@Time : 2020/8/25 11:14
@Author : wangyl
@File : structure.go
@Software: GoLand
*/
package common

import (
	"time"
)

// AuthModel holds information used to authenticate.
type AuthModel struct {
	Username string
	Password string
}

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type BrokenList struct {
	Hosts []struct {
		Hostname string `json:"hostname"`
	} `json:"hosts"`
	Error interface{} `json:"error"`
}

type Msg struct {
	Content string   `json:"content"`
	From    string   `json:"from"`
	Title   string   `json:"title"`
	To      []string `json:"to"`
}

type SingleAlert struct {
	Id       int64             `json:"id"`
	Count    int               `json:"count"`
	Value    float64           `json:"value"`
	Summary  string            `json:"summary"`
	Hostname string            `json:"hostname"`
	Labels   map[string]string `json:"labels"`
}

type Ready2Send struct {
	RuleId int64
	Start  int64
	User   []string
	Alerts []SingleAlert
}

type UserGroup struct {
	Id                    int64
	StartTime             string
	EndTime               string
	Start                 int
	Period                int
	ReversePolishNotation string
	User                  string
	Group                 string
	DutyGroup             string
	Method                string
}

type Alerts []Alert

type Alert struct {
	ActiveAt    time.Time `json:"active_at"`
	Annotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
		RuleId      string `json:"rule_id"`
	} `json:"annotations"`
	FiredAt    time.Time         `json:"fired_at"`
	Labels     map[string]string `json:"labels"`
	LastSentAt time.Time         `json:"last_sent_at"`
	ResolvedAt time.Time         `json:"resolved_at"`
	State      int               `json:"state"`
	ValidUntil time.Time         `json:"valid_until"`
	Value      float64           `json:"value"`
}

type AlertForShow struct {
	Id              int64             `json:"id,omitempty"`
	RuleId          int64             `json:"rule_id"`
	Labels          map[string]string `json:"labels"`
	Value           float64           `json:"value"`
	Count           int               `json:"count"`
	Status          int8              `json:"status"`
	Summary         string            `json:"summary"`
	Description     string            `json:"description"`
	ConfirmedBy     string            `json:"confirmed_by"`
	FiredAt         *time.Time        `json:"fired_at"`
	ConfirmedAt     *time.Time        `json:"confirmed_at"`
	ConfirmedBefore *time.Time        `json:"confirmed_before"`
	ResolvedAt      *time.Time        `json:"resolved_at"`
}

type Confirm struct {
	Duration int
	User     string
	Ids      []int
}

type ValidUserGroup struct {
	User      string
	Group     string
	DutyGroup string
}
