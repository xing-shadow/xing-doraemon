package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

//Plan 告警计划
type Plan struct {
	gorm.Model
	Name      string `gorm:"unique"` //计划名
	StartTime string //告警开始时间
	EndTime   string //告警结束时间

	Method int    //告警方式  0:Hook
	Url    string //告警地址
	Rules  []Rule `gorm:"foreignKey:PlanID"`
}

func (p Plan) TableName() string {
	return "plan"
}

func (p Plan) IsSend(now time.Time) bool {
	data := now.Format("15:04:05")
	return p.StartTime < data && p.EndTime > data
}
