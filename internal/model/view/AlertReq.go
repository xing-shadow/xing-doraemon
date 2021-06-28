package view

import (
	"github.com/prometheus/prometheus/pkg/labels"
	"time"
)

type GetAlertsReq struct {
	Page      uint   `json:"page"`
	PageSize  uint   `json:"page_size"`
	State     string `json:"state"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type ConfirmAlertsReq struct {
	Duration  int    `json:"duration"`
	AlertList []uint `json:"alert_list"`
}

type Alert struct {
	ActiveAt    time.Time     `json:"active_at"`
	Annotations Annotations   `json:"annotations"`
	FiredAt     time.Time     `json:"fired_at"`
	Labels      labels.Labels `json:"labels"`
	LastSentAt  time.Time     `json:"last_sent_at"`
	ResolvedAt  time.Time     `json:"resolved_at"`
	State       int           `json:"state"`
	ValidUntil  time.Time     `json:"valid_until"`
	Value       float64       `json:"value"`
}

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	RuleId      uint   `json:"rule_id"`
}

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}
