package Invoker

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	mysqlDB "xing-doraemon/internal/model/db"
	"xing-doraemon/pkg/Utils"
)

var (
	doraemonMysql *gorm.DB
)

func InitMysqlInvoker() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		mysqlCofig.DBUser,
		mysqlCofig.DBPasswd,
		mysqlCofig.DBDns,
		mysqlCofig.DBName,
		mysqlCofig.DBLoc)

	db, err = gorm.Open(mysqlCofig.DBType, dsn)
	if err != nil {
		if db, err = ensureDatabase(err, dsn, mysqlCofig.DBName); err != nil {
			return
		}
	}
	doraemonMysql = db
	doraemonMysql.SingularTable(true)
	doraemonMysql.LogMode(true)
	doraemonMysql.DB().SetMaxIdleConns(10)
	doraemonMysql.DB().SetMaxOpenConns(100)
	return doraemonMysql, nil
}

func ensureDatabase(errInfo error, dsn string, dbName string) (db *gorm.DB, err error) {
	needInit := false
	switch e := errInfo.(type) {
	case *mysql.MySQLError:
		// MySQL error unkonw database;
		// refer https://dev.mysql.com/doc/refman/5.6/en/error-messages-server.html
		const MysqlErrNum = 1049
		if e.Number == MysqlErrNum {
			needInit = true
			dbForCreateDatabase, err := gorm.Open("mysql", GetTrimDBName(dsn))
			if err != nil {
				return nil, err
			}
			defer dbForCreateDatabase.Close()
			_, err = dbForCreateDatabase.DB().Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;", dbName))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errInfo
		}
	default:
		return nil, errInfo
	}
	// database created, maybe by DBA, but tables not created yet
	if needInit {
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		db.LogMode(true)
		models := []interface{}{
			&mysqlDB.Alert{},
			&mysqlDB.Rule{},
			&mysqlDB.Prom{},
			&mysqlDB.User{},
			&mysqlDB.Plan{},
		}
		db.CreateTable()
		db.SingularTable(true)
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
		db.Create(&mysqlDB.User{
			Name:     "xing",
			Password: Utils.Md5ToHex([]byte("123456")),
		})
	}
	return
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
