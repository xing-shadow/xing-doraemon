/*
 * @Time : 2020/11/16 10:25
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package UserService

import "github.com/jinzhu/gorm"

var opt Option

type Option struct {
	*gorm.DB
}

func Init(option Option) {
	opt = option
}
