/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : plan.go
@Software: GoLand
*/
package model

type Plans struct {
	Id          int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	RuleLabels  string `gorm:"column:rule_labels;size:255" json:"rule_labels"`
	Description string `gorm:"column:description;size:1023" json:"description"`
}

func (*Plans) TableName() string {
	return "plan"
}
