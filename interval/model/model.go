/*
@Time : 2020/8/25 11:24
@Author : wangyl
@File : model.go
@Software: GoLand
*/
package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"xing-doraemon/pkg/setting/alterGateway"
)

type Engine struct {
	dbConfig *alterGateway.Mysql
	db       *gorm.DB
}

func (e *Engine) NewDB(dbConfig *alterGateway.Mysql) *Engine {
	db := new(Engine)
	db.dbConfig = dbConfig
	return db
}

func (e *Engine) InitDbConnection() error {
	needInit := false
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parsetime=true&loc=%s",
		e.dbConfig.DBUser, e.dbConfig.DBPasswd, e.dbConfig.DBTns, e.dbConfig.DBName, e.dbConfig.DBLoc)
	dbConn, err := gorm.Open(e.dbConfig.DBType, dns)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	if err = e.db.DB().PingContext(ctx); err != nil {
		switch e := err.(type) {
		case *mysql.MySQLError:
			// MySQL error unkonw database;
			// refer https://dev.mysql.com/doc/refman/5.6/en/error-messages-server.html
			const MysqlErrNum = 1049
			if e.Number == MysqlErrNum {
				needInit = true
				dbForCreateDatabase, err := sql.Open(DbDriverName, addLocation(dbURL))
				if err != nil {
					return err
				}
				defer dbForCreateDatabase.Close()
				_, err = dbForCreateDatabase.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8 COLLATE utf8_general_ci;", dbName))
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
	e.db = dbConn
	e.db.LogMode(true)
	e.db.DB().SetMaxIdleConns(10)
	e.db.DB().SetMaxOpenConns(100)
	return nil
}
