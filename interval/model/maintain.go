/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : maintain.go
@Software: GoLand
*/
package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"strings"
	"time"
)

type Maintains struct {
	Id        int64  `gorm:"auto" json:"id,omitempty"`
	Flag      bool   `json:"flag"`
	TimeStart string `gorm:"size(15)" json:"time_start"`
	TimeEnd   string `gorm:"size(15)" json:"time_end"`
	Month     int    `json:"month"`
	DayStart  int8   `json:"day_start"`
	DayEnd    int8   `json:"day_end"`
	//Week_start  int8   `json:"week_start"`
	//Week_end    int8   `json:"week_end"`
	//Month_start int8   `json:"month_start"`
	//Month_end   int8   `json:"month_end"`
	Valid *time.Time `json:"valid"`
}

func (Maintains) TableName() string {
	return "maintain"
}

func (m *Maintains) GetAllMaintains(db *gorm.DB) interface{} {
	maintains := []Maintains{}
	db.Table(m.TableName()).Find(&maintains)
	type data struct {
		Id        int64  `json:"id"`
		TimeStart string `json:"time_start"`
		TimeEnd   string `json:"time_end"`
		Month     string `json:"month"`
		DayStart  int8   `json:"day_start"`
		DayEnd    int8   `json:"day_end"`
		Valid     string `json:"valid"`
	}
	res := []data{}
	for _, i := range maintains {
		monthList := []string{}
		for m := 1; m <= 12; m++ {
			if i.Month&int(math.Pow(2, float64(m))) > 0 {
				monthList = append(monthList, strconv.Itoa(m))
			}
		}
		res = append(res, data{
			Id:        i.Id,
			TimeStart: i.TimeStart,
			TimeEnd:   i.TimeEnd,
			Month:     strings.Join(monthList, ","),
			DayStart:  i.DayStart,
			DayEnd:    i.DayEnd,
			Valid:     i.Valid.Format("2006-01-02 15:04:05"),
		})
	}
	return res
}

func (m *Maintains) AddMaintains(db *gorm.DB, hosts string) error {
	tx := db.Begin()
	err := tx.Create(&m).Error
	if err == nil {
		hosts = strings.Replace(hosts, "\r", "", -1)
		var insertErr error
		for _, i := range strings.Split(hosts, string(10)) {
			if i != "" {
				insertErr = tx.Create(&Hosts{Mid: m.Id, Hostname: i}).Error
				if insertErr != nil {
					break
				}
			}
		}
		if insertErr == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	return errors.Wrap(err, "database insert error")
}

func (m *Maintains) UpdateMaintains(db *gorm.DB, hosts string) error {
	tx := db.Begin()
	err := tx.Model(&Maintains{}).Where("id=?", m.Id).Updates(m).Error
	if err == nil {
		deleteErr := tx.Where("mid=?", m.Id).Delete(&Hosts{}).Error
		if deleteErr == nil {
			hosts = strings.Replace(hosts, "\r", "", -1)
			var insertErr error
			for _, i := range strings.Split(hosts, string(10)) {
				if i != "" {
					insertErr = tx.Create(&Hosts{Mid: m.Id, Hostname: i}).Error
					if insertErr != nil {
						break
					}
				}
			}
			if insertErr == nil {
				tx.Commit()
			} else {
				tx.Rollback()
			}
		} else {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	return errors.Wrap(err, "database update error")
}

func (m *Maintains) DeleteMaintains(db *gorm.DB, id string) error {
	tx := db.Begin()
	err := tx.Delete("id = ?", id).Delete(&Maintains{}).Error
	if err == nil {
		err := tx.Delete("mid = ?", id).Delete(&Hosts{}).Error
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
	}
	return errors.Wrap(err, "database delete error")
}
