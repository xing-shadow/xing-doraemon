/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : host.go
@Software: GoLand
*/
package model

type Hosts struct {
	Id       int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Mid      int64  `json:"mid"`
	Hostname string `gorm:"size:255" json:"hostname"`
}

func (*Hosts) TableName() string {
	return "host"
}
