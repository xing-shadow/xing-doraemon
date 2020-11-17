/*
 * @Time : 2020/11/12 11:21
 * @Author : wangyl
 * @File : Init.go
 * @Software: GoLand
 */
package DingTalkService

import (
	"github.com/go-resty/resty/v2"
)

var opt Option
var dingTalkSrv *DingTalkService

type Option struct {
	PushAddr string
	API      *resty.Client
}

func Init(option Option) (err error) {
	opt = option
	opt.API = resty.New()
	dingTalkSrv, err = NewDingTalkService()
	return
}

func PushDingTalkInfo(args interface{}) error {
	return dingTalkSrv.pool.Invoke(args)
}
