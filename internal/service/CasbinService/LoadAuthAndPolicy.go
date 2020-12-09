/*
 * @Time : 2020/11/17 16:15
 * @Author : wangyl
 * @File : LoadAuthAndPolicy.go
 * @Software: GoLand
 */
package CasbinService

import "xing-doraemon/internal/model/db"

func PolicyAuthList() (list []db.CasbinPolicyAuth, err error) {
	err = opt.DB.Find(&list).Error
	return
}

func PolicyGroupList() (list []db.CasbinPolicyGroup, err error) {
	err = opt.DB.Find(&list).Error
	return
}
