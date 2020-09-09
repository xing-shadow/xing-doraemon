/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : host.go
@Software: GoLand
*/
package model

import "github.com/jinzhu/gorm"

type Hosts struct {
	Id       int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Mid      int64  `json:"mid"`
	Hostname string `gorm:"size:255" json:"hostname"`
}

func (Hosts) TableName() string {
	return "host"
}

func (h *Hosts) SetName() {
	h.Mid = 1
}

func (h *Hosts) GetHosts(db *gorm.DB, mid string) []string {
	hosts := []struct {
		Hostname string
	}{}
	db.Table(h.TableName()).Select("hostname").Where("mid=?", mid).Find(&hosts)
	res := []string{}
	for _, i := range hosts {
		res = append(res, i.Hostname)
	}
	return res
}
