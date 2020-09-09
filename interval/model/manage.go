/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : manage.go
@Software: GoLand
*/
package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Manages struct {
	Id          int64  `gorm:"auto" json:"id,omitempty"`
	ServiceName string `gorm:"column:servicename;unique;size:255" json:"servicename"`
	Type        string `gorm:"size:255" json:"type"`
	Status      int8   `gorm:"index" json:"status"`
}

func (Manages) TableName() string {
	return "manage"
}

func (m *Manages) GetAllManage(db *gorm.DB) []Manages {
	manages := []Manages{}
	db.Table(m.TableName()).Find(&manages)
	return manages
}

func (m *Manages) AddManage(db *gorm.DB) error {
	err := db.Create(&m).Error
	return errors.Wrap(err, "database insert error")
}

func (m *Manages) UpdateManage(db *gorm.DB) error {
	err := db.Model(&Manages{}).Where("id=?", m.Id).Updates(m).Error
	return errors.Wrap(err, "database update error")
}

func (p *Manages) DeleteManage(db *gorm.DB, id string) error {
	tx := db.Begin()
	err := tx.Where("id=?", id).Delete(&Manages{}).Error
	if err == nil {
		err := tx.Where("service_id=?", id).Delete(&Configs{}).Error
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "database delete error")
		}
	} else {
		tx.Rollback()
		return errors.Wrap(err, "database delete error")
	}
	tx.Commit()
	return errors.Wrap(err, "database delete error")
}
