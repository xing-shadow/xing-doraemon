package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;size:255" json:"name"`         //用户名
	Password string `gorm:"size:1023" json:"password,omitempty"` //密码
}

func (u User) TableName() string {
	return "user"
}

func (u *User) IsLogin() bool {
	if u.ID > 0 {
		return true
	}
	return false
}
