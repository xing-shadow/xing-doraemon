/*
@Time : 2020/7/20 11:52
@Author : wangyl
@File : Plans.go
@Software: GoLand
*/
package model

type Plans struct {
	Id          int64  `orm:"auto" json:"id,omitempty"`
	RuleLabels  string `orm:"column(rule_labels);size(255)" json:"rule_labels"`
	Description string `orm:"column(description);size(1023)" json:"description"`
}
