/*
 * @Time : 2020/11/12 14:43
 * @Author : wangyl
 * @File : AlertService.go
 * @Software: GoLand
 */
package AlertService

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"os"
	"runtime"
)

type AlertService struct {
	pool *ants.PoolWithFunc
}

func NewAlertService() (obj *AlertService, err error) {
	obj = new(AlertService)
	obj.pool, err = ants.NewPoolWithFunc(1000,
		GetAlertHandle().HandleAlert,
		ants.WithPanicHandler(func(i interface{}) {
			defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 1<<16)
					runtime.Stack(buf, true)
					fmt.Fprintln(os.Stderr, buf)
				}
			}()
		}),
		ants.WithPreAlloc(true),
	)
	return
}

func (s *AlertService) Invoker(task interface{}) error {
	return s.pool.Invoke(task)
}
