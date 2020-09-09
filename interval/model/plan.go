/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : plan.go
@Software: GoLand
*/
package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Plans struct {
	Id          int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	RuleLabels  string `gorm:"column:rule_labels;size:255" json:"rule_labels"`
	Description string `gorm:"column:description;size:1023" json:"description"`
}

func (Plans) TableName() string {
	return "plan"
}

func (plan *Plans) GetAllPlans(db *gorm.DB) []Plans {
	plans := []Plans{}
	db.Find(&plan)
	return plans
}

func (plan *Plans) AddPlan(db *gorm.DB) error {
	err := db.Create(&plan).Error
	return errors.Wrap(err, "database insert error")
}

func (plan *Plans) UpdatePlan(db *gorm.DB) error {
	err := db.Model(&Plans{}).Where("id=?", plan.Id).Updates(plan).Error
	return errors.Wrap(err, "database update error")
}

func (plan *Plans) DeletePlan(db *gorm.DB, id int64) error {
	var rules []struct{ Id int64 }
	tx := db.Begin()
	err := tx.Table(Rules{}.TableName()).Select("id").Where("plan_id = ? LOCK IN SHARE MODE", id).Find(&rules).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		if len(rules) > 0 {
			tx.Commit()
			return fmt.Errorf("cannot delete this plan,it is associated with following rules:%v", rules)
		} else {
			deleErr := tx.Exec("DELETE FROM plan WHERE id = ?", id).Error
			if deleErr != nil {
				deleErr2 := tx.Exec("DELETE FROM plan_receiver WHERE plan_id = ?", id).Error
				if deleErr2 != nil {
					tx.Rollback()
					return errors.Wrap(err, "database delete error")
				}
			} else {
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
