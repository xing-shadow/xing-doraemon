package xtime

import "time"

func ToDuration(val string) time.Duration {
	dst, _ := time.ParseDuration(val)
	return dst
}

func ZeroTime() time.Time {
	timeZero, _ := time.ParseInLocation("2006-01-02 15:04:05", "0001-01-01 00:00:00", time.Local)
	return timeZero
}
