/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : rule.go
@Software: GoLand
*/
package model

type Rules struct {
	Id          int64  `gorm:"column:id;auto" json:"id,omitempty"`
	Expr        string `gorm:"column:expr;size:1023" json:"expr"`
	Op          string `gorm:"column:op;size:31" json:"op"`
	Value       string `gorm:"column:value;size:1023" json:"value"`
	For         string `gorm:"column:for;size:1023" json:"for"`
	Summary     string `gorm:"column:summary;size:1023" json:"summary"`
	Description string `gorm:"column:description;size:1023" json:"description"`
	Prom        *Proms `gorm:"rel:fk" json:"prom_id"`
	Plan        *Plans `gorm:"rel:fk" json:"plan_id"`
	//Labels      []*Labels `gorm:"rel:m2m;rel_through:alert-gateway/models.RuleLabels" json:"omitempty"`
}
