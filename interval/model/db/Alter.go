/*
@Time : 2020/8/25 10:37
@Author : wangyl
@File : Alter.go
@Software: GoLand
*/
package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Alert struct {
	gorm.Model
	Labels      string
	Value       float64
	Count       int
	Status      int8
	Summary     string
	Description string
	Instance    string
	ConfirmedBy string
	FiredAt     *time.Time
	ConfirmedAt *time.Time

	RuleId uint
	Rule   Rule `gorm:"ForeignKey:RuleId"`
}

func (receiver Alert) TableName() string {
	return "alert"
}
