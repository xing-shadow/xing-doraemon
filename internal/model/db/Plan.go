package db

import "github.com/jinzhu/gorm"

//Plan 告警计划
type Plan struct {
	gorm.Model
	Name      string `gorm:"unique"` //计划名
	StartTime string //开始时间
	EndTime   string //结束时间

	Method int    //告警方式  0:Hook
	Url    string //告警地址
	Rules  []Rule `gorm:"foreignKey:PlanID"`
}

func (p Plan) TableName() string {
	return "plan"
}
