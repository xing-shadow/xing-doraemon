package AlertService

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"testing"
	"time"
	"xing-doraemon/configs"
	"xing-doraemon/internal/Invoker"
	"xing-doraemon/internal/model/db"
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
	if err := Invoker.Init(configs.Mysql{
		DBType:   "mysql",
		DBUser:   "root",
		DBPasswd: "123456",
		DBLoc:    "Asia%2FShanghai",
		DBName:   "doraemon",
		DBTns:    "localhost:3306",
	}); err != nil {
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
