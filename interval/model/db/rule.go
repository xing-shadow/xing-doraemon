/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : rule.go
@Software: GoLand
*/
package db

import "github.com/jinzhu/gorm"

type Rule struct {
	gorm.Model
	PlanID      uint
	PromID      uint
	Expr        string
	Op          string
	Value       string
	For         int
	Summary     string
	Description string
	Prom        Prom `gorm:"foreignKey:PromID;"`
	Plan        Plan `gorm:"foreignKey:PlanID"`
}

func (Rule) TableName() string {
	return "rule"
}
