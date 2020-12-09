/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : Users.go
@Software: GoLand
*/
package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;size:255" json:"name"`
	Password string `gorm:"size:1023" json:"password,omitempty"`
}

func (u User) TableName() string {
	return "user"
}
