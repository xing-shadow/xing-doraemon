/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : group.go
@Software: GoLand
*/
package model

type Groups struct {
	Id   int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string `gorm:"unique;size:255" json:"name"`
	User string `gorm:"size:1023" json:"user"`
}

type HttpRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Mobile  string `json:"mobile"`
		Email   string `json:"email"`
		AddTime string `json:"add_time"`
		Account string `json:"account"`
	} `json:"data"`
}

func (*Groups) TableName() string {
	return "group"
}
