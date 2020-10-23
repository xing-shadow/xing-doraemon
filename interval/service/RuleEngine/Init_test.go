/*
 * @Time : 2020/10/23 17:09
 * @Author : wangyl
 * @File : Init_test.go
 * @Software: GoLand
 */
package RuleEngine

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
	"xing-doraemon/pkg/xtime"
)

func TestReloader(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/doraemon?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		t.Fatal(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	Init(Option{
		DB: db,
		Cfg: Config{
			NotifyRetries:      3,
			EvaluationInterval: xtime.ToDuration("15s"),
			ReloadInterval:     xtime.ToDuration("5m"),
		},
	})
}
