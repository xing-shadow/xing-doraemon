/*
@Time : 2020/8/25 10:37
@Author : wangyl
@File : alter.go
@Software: GoLand
*/
package model

import (
	"time"

	"xing-doraemon/pkg/common"
)

type Alerts struct {
	Id              int64      `gorm:"column:id;auto" json:"id,omitempty"`
	Rule            *Rules     `gorm:"rel:fk" json:"rule_id"`
	Labels          string     `gorm:"column:labels;size:4095" json:"labels"`
	Value           float64    `gorm:"column:value" json:"value"`
	Count           int        `json:"count"`
	Status          int8       `gorm:"index" json:"status"`
	Summary         string     `gorm:"column:summary;size:1023" json:"summary"`
	Description     string     `gorm:"column:description;size:1023" json:"description"`
	Hostname        string     `gorm:"column:hostname;size:255" json:"hostname"`
	ConfirmedBy     string     `gorm:"column:confirmed_by;size:1023" json:"confirmed_by"`
	FiredAt         *time.Time `gorm:"type:datetime" json:"fired_at"`
	ConfirmedAt     *time.Time `gorm:"null" json:"confirmed_at"`
	ConfirmedBefore *time.Time `gorm:"null" json:"confirmed_before"`
	ResolvedAt      *time.Time `gorm:"null" json:"resolved_at"`
}

type OneAlert struct {
	ID              int64      `json:"id"`
	RuleID          int64      `json:"rule_id"`
	Value           float64    `json:"value"`
	Status          int8       `json:"status"`
	Count           int        `json:"count"`
	Summary         string     `json:"summary"`
	Description     string     `json:"description"`
	ConfirmedBy     string     `json:"confirmed_by"`
	FiredAt         *time.Time `json:"fired_at"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	ConfirmedBefore *time.Time `json:"confirmed_before"`
	ResolvedAt      *time.Time `json:"resolved_at"`
}

type ShowAlerts struct {
	Alerts []common.AlertForShow `json:"alerts"`
	Total  int64                 `json:"total"`
}

func (Alerts) TableName() string {
	return "alert"
}
