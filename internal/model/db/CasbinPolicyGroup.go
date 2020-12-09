/*
 * @Time : 2020/11/17 16:09
 * @Author : wangyl
 * @File : CasbinPolicyGroup.go
 * @Software: GoLand
 */
package db

import "github.com/jinzhu/gorm"

type CasbinPolicyGroup struct {
	gorm.Model
	UserName  string
	GroupName string
}

func (p CasbinPolicyGroup) TableName() string {
	return "casbin_policy_group"
}
