/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : prom.go
@Software: GoLand
*/
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Proms struct {
	Id   int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string `gorm:"size:1023" json:"name"`
	Url  string `gorm:"size:1023" json:"url"`
}

func (Proms) TableName() string {
	return "prom"
}

func (p *Proms) GetAllProms(db *gorm.DB) []Proms {
	proms := []Proms{}
	db.Table(p.TableName()).Find(&proms)
	return proms
}

func (p *Proms) AddProms(db *gorm.DB) error {
	err := db.Create(&p).Error
	return errors.Wrap(err, "database insert error")
}

func (p *Proms) UpdateProms(db *gorm.DB) error {
	err := db.Model(&Proms{}).Where("id=?", p.Id).Updates(&p).Error
	return errors.Wrap(err, "database update error")
}

func (p *Proms) DeleteProms(db *gorm.DB, id string) error {
	var rules []struct{ Id int64 }
	tx := db.Begin()
	err := tx.Table(Rules{}.TableName()).Select("id").Where("prom_id = ? LOCK IN SHARE MODE", id).Find(&rules).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		if len(rules) > 0 {
			tx.Commit()
			return fmt.Errorf("cannot delete this record,it is associated with following rules:%v", rules)
		} else {
			deleteErr := tx.Where("id=?", id).Delete(&Proms{}).Error
			if deleteErr != nil {
				tx.Rollback()
				return errors.Wrap(err, "database delete error")
			}
		}
	} else {
		tx.Rollback()
		return errors.Wrap(err, "database query error")
	}
	tx.Commit()
	return errors.Wrap(err, "database delete error")
}
