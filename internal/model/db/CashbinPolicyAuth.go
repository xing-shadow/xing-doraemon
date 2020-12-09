/*
 * @Time : 2020/11/17 16:09
 * @Author : wangyl
 * @File : CashbinPolicyAuth.go
 * @Software: GoLand
 */
package db

import "github.com/jinzhu/gorm"

type CasbinPolicyAuth struct {
	gorm.Model
	Sub string `gorm:"not null;" json:"sub"`
	Obj string `gorm:"type:varchar(255);not null" json:"obj"`
	Act string `gorm:"type:varchar(255);not null" json:"act"`
}

func (t CasbinPolicyAuth) TableName() string {
	return "casbin_policy_auth"
}
