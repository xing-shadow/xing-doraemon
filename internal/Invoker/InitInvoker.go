package Invoker

import (
	"github.com/jinzhu/gorm"
	"xing-doraemon/configs"
)

var (
	MysqlInvoker *gorm.DB
	mysqlCofig   configs.Mysql
)

func Init(cfg configs.Mysql) error {
	mysqlCofig = cfg
	if db, err := InitMysqlInvoker(); err != nil {
		return err
	} else {
		MysqlInvoker = db
	}
	return nil
}
