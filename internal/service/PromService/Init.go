package PromService

import "github.com/jinzhu/gorm"

var opt Option

type Option struct {
	*gorm.DB
}

func Init(option Option) {
	opt = option
}
