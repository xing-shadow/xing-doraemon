/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : maintain.go
@Software: GoLand
*/
package model

import (
	"time"
)

type Maintains struct {
	Id        int64  `gorm:"auto" json:"id,omitempty"`
	Flag      bool   `json:"flag"`
	TimeStart string `gorm:"size(15)" json:"time_start"`
	TimeEnd   string `gorm:"size(15)" json:"time_end"`
	Month     int    `json:"month"`
	DayStart  int8   `json:"day_start"`
	DayEnd    int8   `json:"day_end"`
	//Week_start  int8   `json:"week_start"`
	//Week_end    int8   `json:"week_end"`
	//Month_start int8   `json:"month_start"`
	//Month_end   int8   `json:"month_end"`
	Valid *time.Time `json:"valid"`
}

func (Maintains) TableName() string {
	return "maintain"
}
