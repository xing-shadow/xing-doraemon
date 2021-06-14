package db

import "github.com/jinzhu/gorm"

//Prom Prometheus 地址
type Prom struct {
	gorm.Model
	Name string `gorm:"size:1023,unique" json:"name"` //名称
	Url  string `gorm:"size:1023" json:"url"`         //Prometheus地址
}

func (Prom) TableName() string {
	return "prom"
}
