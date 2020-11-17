/*
 * @Time : 2020/10/19 15:53
 * @Author : wangyl
 * @File : Plan.go
 * @Software: GoLand
 */
package db

import "github.com/jinzhu/gorm"

type Plan struct {
	gorm.Model
	Name       string `gorm:"unique;"`
	StartTime  string
	EndTime    string
	Expression string
	Period     int
	User       string
	Rules      []Rule `gorm:"foreignKey:PlanID"`
}

func (p Plan) TableName() string {
	return "plan"
}