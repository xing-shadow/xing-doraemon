package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	Labels      string  //标签
	Value       float64 //告警值
	Status      int8    //告警状态 0:正常 1:挂起 2:触发
	Summary     string  //告警简介
	Description string  //描述
	Instance    string  //节点
	ConfirmedBy string  //确认人

	FiredAt         *time.Time //第一次触发时间
	ConfirmedAt     *time.Time //告警确认时间
	ConfirmedBefore *time.Time //告警确认时长
	ResolvedAt      *time.Time //告警解除时间

	RuleId uint
	Rule   Rule `gorm:"ForeignKey:RuleId"`
}

func (receiver Alert) TableName() string {
	return "alert"
}
