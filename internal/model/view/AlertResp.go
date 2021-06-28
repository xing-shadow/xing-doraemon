package view

import (
	"time"
)

type GetAlertResp struct {
	Pagination `json:"pagination"`
	Alerts     Alerts `json:"list"`
}

type Alerts []AlertItem

type AlertItem struct {
	ID              uint       `json:"id"`
	RuleId          uint       `json:"rule_id"`
	Labels          string     `json:"labels"`
	Value           float64    `json:"value"`
	Status          int        `json:"status"`
	Summary         string     `json:"summary"`
	Instance        string     `json:"instance"`
	Description     string     `json:"description"`
	ConfirmedBy     string     `json:"confirmed_by"`
	FiredAt         *time.Time `json:"fired_at"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	ConfirmedBefore *time.Time `json:"confirmed_before"`
	ResolvedAt      *time.Time `json:"resolved_at"`
}
