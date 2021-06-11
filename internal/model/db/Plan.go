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
