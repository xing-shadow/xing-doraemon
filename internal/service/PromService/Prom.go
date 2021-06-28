package PromService

import (
	"errors"

	"github.com/jinzhu/gorm"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/model/view"
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
	err = opt.DB.Table(db.Prom{}.TableName()).Select("id, name, url").Offset(offset).Count(&count).Limit(pageSize).Find(&proms).Error
	if err != nil {
		return
	}
	resp.CurrentPage = page
	resp.Total = count
	resp.PageSize = pageSize
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

	err = opt.DB.Select("id, name, url").Where("id=?", req.ID).First(&prom).Error
	if err != nil {
		return
	}
	resp.Url = prom.Url
	resp.Name = prom.Name
	resp.ID = prom.ID
	return
}

func GetPromAllName() (resp []string, err error) {
	var results []db.Prom
	err = opt.DB.Select("name").Find(&results).Error
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
	prom.Name = req.Name
	prom.Url = req.Url
	err = opt.DB.Save(&prom).Error
	return
}

func ModifyProm(req view.ModifyProm) (err error) {
	var prom db.Prom
	err = opt.DB.Where("id=?", req.ID).First(&prom).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("该数据源不存在")
			return
		}
		return
	}
	prom.Name = req.Name
	prom.Url = req.Url
	err = opt.DB.Save(&prom).Error
	return
}

func DeleteProm(req view.DeleteProm) (err error) {
	var prom db.Prom
	err = opt.DB.Where("id = ?", req.ID).Delete(&prom).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("该记录不存在")
			return
		}
		return err
	}
	return nil
}
