package RuleService

import "github.com/jinzhu/gorm"

var opt Option

type Option struct {
	DB *gorm.DB
}

func Init(option Option) {
	opt = option
}
