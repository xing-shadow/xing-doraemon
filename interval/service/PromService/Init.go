/*
 * @Time : 2020/10/22 15:51
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package PromService

import "github.com/jinzhu/gorm"

var opt Option

type Option struct {
	*gorm.DB
}

func Init(option Option) {
	opt = option
}
