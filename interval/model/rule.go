/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : rule.go
@Software: GoLand
*/
package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"xing-doraemon/gobal"
)

type Rules struct {
	Id          int64  `gorm:"column:id;auto" json:"id,omitempty"`
	Expr        string `gorm:"column:expr;size:1023" json:"expr"`
	Op          string `gorm:"column:op;size:31" json:"op"`
	Value       string `gorm:"column:value;size:1023" json:"value"`
	For         string `gorm:"column:for;size:1023" json:"for"`
	Summary     string `gorm:"column:summary;size:1023" json:"summary"`
	Description string `gorm:"column:description;size:1023" json:"description"`
	Prom        *Proms `gorm:"rel:fk" json:"prom_id"`
	Plan        *Plans `gorm:"rel:fk" json:"plan_id"`
	//Labels      []*Labels `gorm:"rel:m2m;rel_through:alert-gateway/models.RuleLabels" json:"omitempty"`
}

func (Rules) TableName() string {
	return "rule"
}

func (r *Rules) Get(db *gorm.DB, prom string, id string) ([]Rules, error) {
	var result []Rules
	var err error
	if prom != "" {
		err = db.Table(r.TableName()).Where("prom_id = ?", prom).Find(&result).Error
	} else if id != "" {
		err = db.Table(r.TableName()).Where("id = ?", prom).Find(&result).Error
	} else {
		err = db.Table(r.TableName()).Find(&result).Error
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}
	return result, nil
}

func (r *Rules) InsertRule(db *gorm.DB) error {
	logger := gobal.GetLogger()
	var prom []struct{ PromId int64 }
	var plan []struct{ PlanId int64 }
	tx := db.Begin()
	err := tx.Table(Proms{}.TableName()).Select("id").Where("id=? LOCK IN SHARE MODE", r.Prom.Id).Find(&prom).Error
	if err == nil && len(prom) > 0 {
		err = tx.Table(Plans{}.TableName()).Select("id").Where("id=? LOCK IN SHARE MODE", r.Plan.Id).Find(&plan).Error
		if err == nil && len(plan) > 0 {
			err = tx.Table(r.TableName()).Create(r).Error
			if err != nil {
				logger.Errorf("Insert rule error:%v", err)
				tx.Rollback()
				return errors.Wrap(err, "database insert error")
			}
			tx.Commit()
			return nil
		} else {
			tx.Rollback()
			logger.Errorf("The plan_id %s is invalid", r.Plan.Id)
			return fmt.Errorf("invalid plan_id %v", r.Plan.Id)
		}
	} else {
		tx.Rollback()
		logger.Errorf("The prom_id %s is invalid", r.Prom.Id)
		return fmt.Errorf("invalid prom_id %v", r.Prom.Id)
	}
}

func (r *Rules) UpdateRule(db *gorm.DB) error {
	logger := gobal.GetLogger()
	var prom []struct{ PromId int64 }
	var plan []struct{ PlanId int64 }
	tx := db.Begin()
	err := tx.Table(Proms{}.TableName()).Select("id").Where("id=? LOCK IN SHARE MODE", r.Prom.Id).Find(&prom).Error
	if err == nil && len(prom) > 0 {
		err = tx.Table(Plans{}.TableName()).Select("id").Where("id=? LOCK IN SHARE MODE", r.Plan.Id).Find(&plan).Error
		//fmt.Println(plan)
		if err == nil && len(plan) > 0 {
			err = tx.Model(&Rules{}).Where("id=?", r.Id).Update(r).Error
			if err != nil {
				logger.Errorf("update rule error:%v", err)
				tx.Rollback()
				return errors.Wrap(err, "database update error")
			}
			tx.Commit()
		} else {
			tx.Rollback()
			logger.Errorf("The plan_id %s is invalid", r.Plan.Id)
			return fmt.Errorf("invalid plan_id %v", r.Plan.Id)
		}
	} else {
		tx.Rollback()
		logger.Errorf("The prom_id %s is invalid", r.Prom.Id)
		return fmt.Errorf("invalid prom_id %v", r.Prom.Id)
	}
	tx.Commit()
	return errors.Wrap(err, "database update error")
}

func (r *Rules) DeleteRule(db *gorm.DB, Id string) error {
	err := db.Where("id=?", Id).Delete(&Rules{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
