/*
 * @Time : 2020/10/23 14:56
 * @Author : wangyl
 * @File : ParseTime.go
 * @Software: GoLand
 */
package xtime

import "time"

func ToDuration(val string) time.Duration {
	dst, _ := time.ParseDuration(val)
	return dst
}
