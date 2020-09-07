/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : prom.go
@Software: GoLand
*/
package model

type Proms struct {
	Id   int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string `gorm:"size:1023" json:"name"`
	Url  string `gorm:"size:1023" json:"url"`
}

func (Proms) TableName() string {
	return "prom"
}
