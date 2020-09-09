/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : config.go
@Software: GoLand
*/
package model

import (
	xerrors "errors"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var ErrNoService = errors.New("add config failed:the service is not exist,please refresh")

type Configs struct {
	Id        int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	ServiceId int64  `json:"serviceid"`
	Idc       string `gorm:"size:255" json:"idc"`
	Proto     string `gorm:"size:255" json:"proto"`
	Auto      string `gorm:"size:255" json:"auto"`
	Port      int    `json:"port"`
	Metric    string `gorm:"size:255" json:"metric"`
}

func (*Configs) TableName() string {
	return "config"
}

func (c *Configs) GetAllConfig(db *gorm.DB, idc string) []Configs {
	configs := []Configs{}
	if idc != "" {
		db.Table(c.TableName()).Where("idc=?", idc).Find(&configs)
	} else {
		db.Table(c.TableName()).Find(&configs)
	}
	return configs
}

func (c *Configs) AddConfig(db *gorm.DB) error {
	var rows []struct{ Id int64 }
	tx := db.Begin()
	err := tx.Table(Manages{}.TableName()).Select("id").Where("id=? LOCK IN SHARE MODE", c.ServiceId).Find(&rows).Error
	if err == nil || xerrors.Is(err, gorm.ErrRecordNotFound) {
		if len(rows) > 0 {
			err = tx.Create(&c).Error
			if err == nil {
				tx.Commit()
			} else {
				tx.Rollback()
			}
		} else {
			tx.Commit()
			return ErrNoService
		}
	} else {
		tx.Rollback()
		return errors.Wrap(err, "database query error")
	}
	return errors.Wrap(err, "database insert error")
}

func (c *Configs) UpdateConfig(db *gorm.DB) error {
	err := db.Model(&Configs{}).Where("id=?", c.Id).Updates(map[string]interface{}{
		"idc":    c.Idc,
		"proto":  c.Proto,
		"metric": c.Metric,
		"auto":   c.Auto,
		"port":   c.Proto,
	}).Error
	return errors.Wrap(err, "database update error")
}

func (c *Configs) DeleteConfig(db *gorm.DB, id string) error {
	err := db.Where("id=?", c.Id).Delete(&Configs{}).Error
	return errors.Wrap(err, "database delete error")
}
