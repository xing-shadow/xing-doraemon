package Invoker

import (
	"testing"
	"time"
	"xing-doraemon/configs"
	"xing-doraemon/internal/model/db"
)

func TestConnMysql(t *testing.T) {
	err := Init(configs.Mysql{
		DBType:    "mysql",
		DBUser:    "root",
		DBPasswd:  "123456",
		DBLoc:     "Asia%2FShanghai",
		DBConnTTL: 5,
		DBName:    "doraemon",
		DBDns:     "localhost:3306",
	})
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	var alert = db.Alert{
		FiredAt:         &now,
		ConfirmedAt:     &now,
		ConfirmedBefore: &now,
		ResolvedAt:      nil,
	}
	err = doraemonMysql.Save(&alert).Error
	if err != nil {
		t.Fatal(err)
	}
}
