package AlertService

import (
	"github.com/jinzhu/gorm"
)

var opt Option

type Option struct {
	DB *gorm.DB
}

func Init(option Option) (err error) {
	opt = option
	//go alertInspection()
	return
}
