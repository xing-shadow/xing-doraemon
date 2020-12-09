/*
 * @Time : 2020/11/12 14:43
 * @Author : wangyl
 * @File : AlertService_test.go
 * @Software: GoLand
 */
package AlertService

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"testing"
	"time"
	"xing-doraemon/global"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model/db"
)

func TestA(t *testing.T) {
	p, _ := ants.NewPool(1000, ants.WithPanicHandler(func(i interface{}) {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
	}))
	for {
		err := p.Submit(func() {
			fmt.Println("hello world")
		})
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(time.Second)
	}
	select {}
}

func TestConfirmAlertList(t *testing.T) {
	p, _ := ants.NewPoolWithFunc(1000, func(i interface{}) {
		fmt.Println("heihei")
		time.Sleep(time.Second * 5)
		fmt.Println("heihei1")
	}, ants.WithPanicHandler(func(i interface{}) {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
	}), ants.WithExpiryDuration(time.Second*1))
	p.Invoke(1)
	for {
		fmt.Println(p.Free())
		time.Sleep(time.Second)
	}
}

func TestQueryAlert(t *testing.T) {
	global.GetAlterGatewayConfig().Mysql.DBType = "mysql"
	global.GetAlterGatewayConfig().Mysql.DBUser = "root"
	global.GetAlterGatewayConfig().Mysql.DBPasswd = "123456"
	global.GetAlterGatewayConfig().Mysql.DBTns = "localhost:3306"
	global.GetAlterGatewayConfig().Mysql.DBName = "doraemon" //doraemon
	global.GetAlterGatewayConfig().Mysql.DBLoc = "Asia%2FShanghai"
	if err := Invoker.Init(); err != nil {
		t.Fatal(err)
	}
	opt.DB = Invoker.MysqlInvoker
	var alerts []db.Alert
	err := opt.DB.Where("status=? and confirmed_at is not null and ? > confirmed_at ", 2, time.Now()).Find(&alerts).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(alerts)
}

func TestB(t *testing.T) {
	fmt.Println(time.Now().Add(-2 * time.Hour).Format(time.RFC3339))
}
