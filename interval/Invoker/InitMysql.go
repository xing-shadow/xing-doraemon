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
)

var (
	doraemonMysql *gorm.DB
)

func GetDB() *gorm.DB {
	return doraemonMysql
}

func Init() error {
	if err := InitDB(); err != nil {
		return err
	}
	return nil
}

func InitDB() error {
	cfg := gobal.GetAlterGatewayConfig().Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		cfg.DBUser,
		cfg.DBPasswd,
		cfg.DBTns,
		cfg.DBName,
		cfg.DBLoc)
	db, err := gorm.Open(cfg.DBType, dsn)
	if err != nil {
		return err
	}
	if err := ensureDatabase(db, dsn, cfg.DBName); err != nil {
		return err
	}
	doraemonMysql = db
	doraemonMysql.SingularTable(true)
	doraemonMysql.LogMode(true)
	doraemonMysql.DB().SetMaxIdleConns(10)
	doraemonMysql.DB().SetMaxOpenConns(100)
	return nil
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
	if !needInit {
		sql := fmt.Sprintf("show tables from %s;", dbName)
		if rows, err := db.DB().Query(sql); err == nil && !rows.Next() {
			needInit = true
		}
	}
	gobal.GetLogger().Debugf("Initialize database connection: %s", dsn)
	if needInit {
		for _, insertSql := range InitialData {
			_, err := db.DB().Exec(insertSql)
			if err != nil {
				return err
			}
		}
	}
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
