/*
 * @Time : 2020/10/23 15:58
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package AlertService

import "github.com/jinzhu/gorm"

var opt Option

type Option struct {
	DB *gorm.DB
}

func Init(option Option) {
	opt = option
}
