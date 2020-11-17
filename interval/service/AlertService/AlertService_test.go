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
