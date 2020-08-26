/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : config.go
@Software: GoLand
*/
package model

import (
	"errors"
)

var ErrNoService = errors.New("add config failed:the service is not exist,please refresh")

type Configs struct {
	Id        int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	ServiceId int64  `json:"serviceid"`
	Idc       string `gorm:"size:255" json:"idc"`
	Proto     string `gorm:"size:255" json:"proto"`
	Auto      string `gorm:"size:255" json:"auto"`
	Port      int    `json:"port"`
	Metric    string `gorm:"size:255" json:"metric"`
}

func (*Configs) TableName() string {
	return "config"
}
