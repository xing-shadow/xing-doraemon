/*
@Time : 2020/9/3 16:33
@Author : wangyl
@File : Init.go
@Software: GoLand
*/
package Invoker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xing-doraemon/gobal"
	mysqlDB "xing-doraemon/interval/model/db"
)

var (
	doraemonMysql *gorm.DB
)

func InitMysqlInvoker() (*gorm.DB, error) {
	cfg := gobal.GetAlterGatewayConfig().Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		cfg.DBUser,
		cfg.DBPasswd,
		cfg.DBTns,
		cfg.DBName,
		cfg.DBLoc)
	db, err := gorm.Open(cfg.DBType, dsn)
	if err != nil {
		return nil, err
	}
	if err := ensureDatabase(db, dsn, cfg.DBName); err != nil {
		return nil, err
	}
	doraemonMysql = db
	doraemonMysql.SingularTable(true)
	doraemonMysql.LogMode(true)
	doraemonMysql.DB().SetMaxIdleConns(10)
	doraemonMysql.DB().SetMaxOpenConns(100)
	return doraemonMysql, nil
}

func ensureDatabase(db *gorm.DB, dsn string, dbName string) error {
	needInit := false
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	if err := db.DB().PingContext(ctx); err != nil {
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
	}
	// database created, maybe by DBA, but tables not created yet
	if needInit {
		models := []interface{}{
			&mysqlDB.Alert{},
			&mysqlDB.Rule{},
			&mysqlDB.Prom{},
			&mysqlDB.User{},
			&mysqlDB.Plan{},
		}
		db.SingularTable(true)
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
		gobal.GetLogger().Info("create tables ok")
	}
	gobal.GetLogger().Debugf("Initialize database connection: %s", dsn)
	return nil
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
