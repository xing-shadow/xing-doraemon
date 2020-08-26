/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : manage.go
@Software: GoLand
*/
package model

type Manages struct {
	Id          int64  `gorm:"auto" json:"id,omitempty"`
	ServiceName string `gorm:"column:servicename;unique;size:255" json:"servicename"`
	Type        string `gorm:"size:255" json:"type"`
	Status      int8   `gorm:"index" json:"status"`
}

func (*Manages) TableName() string {
	return "manage"
}
