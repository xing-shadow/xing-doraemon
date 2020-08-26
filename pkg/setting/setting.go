/*
@Time : 2020/8/24 17:35
@Author : wangyl
@File : setting.go
@Software: GoLand
*/
package setting

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configPath string) (*Setting, error) {
	dir := filepath.Dir(configPath)
	parts := strings.Split(filepath.Base(configPath), ".")
	var filename, ext string
	if len(parts) == 2 {
		filename = parts[0]
		ext = parts[1]
	} else {
		return nil, errors.New("invalid config path")
	}
	var setting = &Setting{}
	setting.vp = viper.New()
	setting.vp.SetConfigName(filename)
	setting.vp.AddConfigPath(dir)
	setting.vp.SetConfigType(ext)
	if err := setting.vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return setting, nil
}

func (v *Setting) ReadStruct(val interface{}) error {
	if err := v.vp.Unmarshal(val); err != nil {
		return err
	} else {
		return nil
	}
}

func (v *Setting) ReadSection(key string) interface{} {
	return v.vp.Get(key)
}
