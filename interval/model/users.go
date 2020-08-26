/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : users.go
@Software: GoLand
*/
package model

type Users struct {
	Id       int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name     string `gorm:"unique;size:255" json:"name"`
	Password string `gorm:"size:1023" json:"password,omitempty"`
}

func (u *Users) TableName() string {
	return "users"
}
