/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : prom.go
@Software: GoLand
*/
package db

import "github.com/jinzhu/gorm"

type Prom struct {
	gorm.Model
	Name string `gorm:"size:1023,unique" json:"name"`
	Url  string `gorm:"size:1023" json:"url"`
}

func (Prom) TableName() string {
	return "prom"
}
