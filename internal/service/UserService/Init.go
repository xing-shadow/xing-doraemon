/*
 * @Time : 2020/11/16 10:25
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package UserService

import (
	"github.com/jinzhu/gorm"
	"xing-doraemon/configs"
)

var opt Option
var UserSession userSession

type Option struct {
	DB      *gorm.DB
	SessCfg configs.Session
}

func Init(option Option) {
	opt = option
	InitUser()
	initUserSession()
}
