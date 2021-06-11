package DingTalkService

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"os"
	"runtime"
)

type DingTalkService struct {
	pool *ants.PoolWithFunc
}

func NewDingTalkService() (obj *DingTalkService, err error) {
	obj = new(DingTalkService)
	obj.pool, err = ants.NewPoolWithFunc(1000,
		HandleDingTalkInfoFunc,
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
