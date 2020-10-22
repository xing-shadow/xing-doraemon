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
		pageSize = int(req.PageSize)
	} else {
		pageSize = int(req.PageSize)
	}
	offset = (page - 1) * pageSize
	err = opt.DB.Offset(offset).Limit(pageSize).Find(&proms).Error
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
			Name: prom.Name,
			Url:  prom.Url,
		})
	}
	return
}

func GetAllProms() (resp view.PromList, err error) {
	var proms []db.Prom
	var count int
	err = opt.DB.Find(&proms).Error
	if err != nil {
		return
	}
	err = opt.DB.Model(&db.Prom{}).Count(&count).Error
	if err != nil {
		return
	}
	resp.Total = count
	for _, prom := range proms {
		resp.PromList = append(resp.PromList, view.PromItem{
			Name: prom.Name,
			Url:  prom.Url,
		})
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
