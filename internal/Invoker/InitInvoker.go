/*
 * @Time : 2020/10/19 16:38
 * @Author : wangyl
 * @File : InitInvoker.go
 * @Software: GoLand
 */
package Invoker

import "github.com/jinzhu/gorm"

var (
	MysqlInvoker *gorm.DB
)

func Init() error {
	if db, err := InitMysqlInvoker(); err != nil {
		return err
	} else {
		MysqlInvoker = db
	}
	return nil
}
