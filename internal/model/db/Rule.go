package db

import "github.com/jinzhu/gorm"

//Rule 告警规则
type Rule struct {
	gorm.Model
	PlanID      uint   `gorm:"unique"` // 计划id
	PromID      uint   `gorm:"unique"` //prom id
	Expr        string //PromQL
	Op          string //算数运算符
	Value       string //阈值
	For         int    //持续时间
	Summary     string //简介
	Description string //描述
	Prom        Prom   `gorm:"foreignKey:PromID;"`
	Plan        Plan   `gorm:"foreignKey:PlanID"`
}

func (Rule) TableName() string {
	return "rule"
}
