/*
 * @Time : 2020/10/22 15:53
 * @Author : wangyl
 * @File : Prom.go
 * @Software: GoLand
 */
package PromService

import (
	"errors"
	"github.com/jinzhu/gorm"
	"xing-doraemon/interval/model/db"
	"xing-doraemon/interval/model/view"
)

func GetPromPagination(req view.GetProms) (resp view.PromList, err error) {
	var page, pageSize, offset int
	var proms []db.Prom
	var count int
	if req.Page == 0 {
		page = 1
	} else {
		page = int(req.Page)
	}
	if req.PageSize == 0 {
		pageSize = 1000
	} else {
		pageSize = int(req.PageSize)
	}
	offset = (page - 1) * pageSize
	if req.Name != "" && req.Url != "" {
		err = opt.DB.Where("name LIKE ? AND url LIKE ? ",
			"%"+req.Name+"%",
			"%"+req.Url+"%").Offset(offset).Limit(pageSize).Find(&proms).Error
	} else if req.Name != "" && req.Url == "" {
		err = opt.DB.Where("name LIKE ?",
			"%"+req.Name+"%").Offset(offset).Limit(pageSize).Find(&proms).Error
	} else if req.Name == "" && req.Url != "" {
		err = opt.DB.Where("url LIKE ? ",
			"%"+req.Url+"%").Offset(offset).Limit(pageSize).Find(&proms).Error
	} else {
		err = opt.DB.Offset(offset).Limit(pageSize).Find(&proms).Error
	}
	if err != nil {
		return
	}
	err = opt.DB.Model(&db.Prom{}).Count(&count).Error
	if err != nil {
		return
	}
	resp.CurrentPage = page
	resp.Total = count
	for _, prom := range proms {
		resp.PromList = append(resp.PromList, view.PromItem{
			ID:   prom.ID,
			Name: prom.Name,
			Url:  prom.Url,
		})
	}
	return
}

func GetProm(req view.GetProm) (resp view.PromItem, err error) {
	var prom db.Prom

	err = opt.DB.Where("id=?", req.ID).First(&prom).Error
	if err != nil {
		return
	}
	resp.Url = prom.Url
	resp.Name = prom.Name
	resp.ID = prom.ID
	return
}

func GetPromAllName() (resp []string, err error) {
	var results = []struct {
		Name string `gorm:"column:name;"`
	}{}
	err = opt.DB.Table(db.Prom{}.TableName()).Select("name").Find(&results).Error
	if err != nil {
		return
	}
	for _, result := range results {
		resp = append(resp, result.Name)
	}
	return
}

func CreateProms(req view.CreateProm) (err error) {
	var prom db.Prom
	err = opt.DB.Where("name=? and url=?", req.Name, req.Url).First(&prom).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	} else {
		return errors.New("this prom exist")
	}
	err = opt.DB.Create(&db.Prom{
		Name: req.Name,
		Url:  req.Url,
	}).Error
	return
}

func ModifyProm(req view.ModifyProm) (err error) {
	var prom db.Prom
	err = opt.DB.Where("id=?", req.ID).First(&prom).Error
	if err != nil {
		return err
	}
	err = opt.DB.Model(&db.Prom{}).Where("id=?", req.ID).Updates(&db.Prom{
		Name: req.Name,
		Url:  req.Url,
	}).Error
	return
}

func DeleteProm(req view.DeleteProm) (err error) {
	var prom db.Prom
	err = opt.DB.Where("id = ?", req.ID).Delete(&prom).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return err
	}
	return nil
}
