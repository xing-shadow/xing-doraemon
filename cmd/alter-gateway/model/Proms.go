/*
@Time : 2020/7/20 11:52
@Author : wangyl
@File : Proms.go
@Software: GoLand
*/
package model

type Proms struct {
	Id   int64  `orm:"auto" json:"id,omitempty" bson:"Id"`
	Name string `orm:"size(1023)" json:"name" bson:"Name"`
	Url  string `orm:"size(1023)" json:"url" bson:"Url"`
}
