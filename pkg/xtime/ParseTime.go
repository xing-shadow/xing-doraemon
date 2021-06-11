package xtime

import "time"

func ToDuration(val string) time.Duration {
	dst, _ := time.ParseDuration(val)
	return dst
}
