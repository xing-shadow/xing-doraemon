/*
 * @Time : 2020/10/19 15:53
 * @Author : wangyl
 * @File : plan.go
 * @Software: GoLand
 */
package db

import "github.com/jinzhu/gorm"

type Plan struct {
	gorm.Model
	StartTime  string
	EndTime    string
	Period     int
	Expression string
	User       string
	Rules      []Rule `gorm:"foreignKey:PlanID"`
}

func (p Plan) TableName() string {
	return "plan"
}
