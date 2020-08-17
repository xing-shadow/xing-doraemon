/*
@Time : 2020/7/20 11:42
@Author : wangyl
@File : Usres.go
@Software: GoLand
*/
package model

type Users struct {
	Id       int64  `orm:"auto" json:"id,omitempty" bson:"Id"`
	Name     string `orm:"unique;size(255)" json:"name" bson:"Name"`
	Password string `orm:"size(1023)" json:"password,omitempty" bson:"Password"`
}

func (*Users) TableName() string {
	return "Users"
}
