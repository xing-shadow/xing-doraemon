/*
@Time : 2020/9/3 16:33
@Author : wangyl
@File : Init.go
@Software: GoLand
*/
package Invoker

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"xing-doraemon/global"
	mysqlDB "xing-doraemon/interval/model/db"
	"xing-doraemon/pkg/Utils"
)

var (
	doraemonMysql *gorm.DB
)

func InitMysqlInvoker() (*gorm.DB, error) {
	cfg := global.GetAlterGatewayConfig().Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		cfg.DBUser,
		cfg.DBPasswd,
		cfg.DBTns,
		cfg.DBName,
		cfg.DBLoc)
	db, err := gorm.Open(cfg.DBType, dsn)
	if err != nil {
		fmt.Println(err)
		if err := ensureDatabase(err, dsn, cfg.DBName); err != nil {
			return nil, err
		}
	}
	ensureTables(db)
	doraemonMysql = db
	doraemonMysql.SingularTable(true)
	doraemonMysql.LogMode(true)
	doraemonMysql.DB().SetMaxIdleConns(10)
	doraemonMysql.DB().SetMaxOpenConns(100)
	return doraemonMysql, nil
}

func ensureDatabase(err error, dsn string, dbName string) error {
	needInit := false
	switch e := err.(type) {
	case *mysql.MySQLError:
		// MySQL error unkonw database;
		// refer https://dev.mysql.com/doc/refman/5.6/en/error-messages-server.html
		const MysqlErrNum = 1049
		if e.Number == MysqlErrNum {
			needInit = true
			dbForCreateDatabase, err := gorm.Open("mysql", GetTrimDBName(dsn))
			if err != nil {
				return err
			}
			defer dbForCreateDatabase.Close()
			_, err = dbForCreateDatabase.DB().Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8 COLLATE utf8_general_ci;", dbName))
			if err != nil {
				return err
			}
		} else {
			return err
		}
	default:
		return err
	}
	// database created, maybe by DBA, but tables not created yet
	if needInit {
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			return err
		}
		models := []interface{}{
			&mysqlDB.Alert{},
			&mysqlDB.Rule{},
			&mysqlDB.Prom{},
			&mysqlDB.User{},
			&mysqlDB.Plan{},
			&mysqlDB.CasbinPolicyGroup{},
			&mysqlDB.CasbinPolicyAuth{},
		}
		db.Create(&mysqlDB.User{
			Name:     "xing",
			Password: Utils.Md5ToHex([]byte("123456")),
		})
		db.Create(&mysqlDB.CasbinPolicyGroup{
			UserName:  "xing",
			GroupName: "admin",
		})
		db.SingularTable(true)
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
		global.GetLogger().Info("create tables ok")
	}
	global.GetLogger().Debugf("Initialize database connection: %s", dsn)
	return nil
}

func ensureTables(db *gorm.DB) {
	if ok := db.HasTable(&mysqlDB.Alert{}); !ok {
		db.CreateTable(&mysqlDB.Alert{})
	}
	if ok := db.HasTable(&mysqlDB.Rule{}); !ok {
		db.CreateTable(&mysqlDB.Rule{})
	}
	if ok := db.HasTable(&mysqlDB.Prom{}); !ok {
		db.CreateTable(&mysqlDB.Prom{})
	}
	if ok := db.HasTable(&mysqlDB.User{}); !ok {
		db.CreateTable(&mysqlDB.User{})
		db.Create(&mysqlDB.User{
			Name:     "xing",
			Password: Utils.Md5ToHex([]byte("123456")),
		})
	}

	if ok := db.HasTable(&mysqlDB.Plan{}); !ok {
		db.CreateTable(&mysqlDB.Plan{})
	}
	if ok := db.HasTable(&mysqlDB.CasbinPolicyAuth{}); !ok {
		db.CreateTable(&mysqlDB.CasbinPolicyAuth{})
	}
	if ok := db.HasTable(&mysqlDB.CasbinPolicyGroup{}); !ok {
		db.CreateTable(&mysqlDB.CasbinPolicyGroup{})
		db.Create(&mysqlDB.CasbinPolicyGroup{
			UserName:  "xing",
			GroupName: "admin",
		})
	}
}

func GetTrimDBName(dsn string) string {
	parts := strings.Split(dsn, "/")
	if len(parts) != 2 {
		return dsn
	}
	subparts := strings.Split(parts[1], "?")
	if len(subparts) != 2 {
		return dsn
	}
	return parts[0] + "/?" + subparts[1]
}
