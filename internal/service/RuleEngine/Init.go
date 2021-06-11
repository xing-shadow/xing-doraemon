package RuleEngine

import "github.com/jinzhu/gorm"

var opt Option

type Option struct {
	*gorm.DB
	Cfg Config
}

func Init(option Option) {
	opt = option

	reloader := NewReloader(opt.Cfg)

	reloader.Start()
}
