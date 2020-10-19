/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : users.go
@Software: GoLand
*/
package db

type User struct {
	Id       int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name     string `gorm:"unique;size:255" json:"name"`
	Password string `gorm:"size:1023" json:"password,omitempty"`
}

func (u User) TableName() string {
	return "user"
}
