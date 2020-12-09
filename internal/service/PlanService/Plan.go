/*
 * @Time : 2020/10/22 14:53
 * @Author : wangyl
 * @File : Plan.go
 * @Software: GoLand
 */
package PlanService

import (
	"errors"
	"github.com/jinzhu/gorm"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/model/view"
)

func GetPlan(req view.GetPlan) (resp view.PlanItem, err error) {
	var plan db.Plan
	err = opt.DB.Where("id=?", req.Id).First(&plan).Error
	if err != nil {
		return
	}
	resp = view.PlanItem{
		Id:         plan.ID,
		Name:       plan.Name,
		Expression: plan.Expression,
		StartTime:  plan.StartTime,
		EndTime:    plan.EndTime,
		Period:     plan.Period,
	}
	return
}

func GetPlanAllName() (resp []string, err error) {
	var results = []struct {
		Name string `gorm:"column:name;"`
	}{}
	err = opt.DB.Table(db.Plan{}.TableName()).Select("name").Where("deleted_at is NULL").Find(&results).Error
	if err != nil {
		return
	}
	for _, result := range results {
		resp = append(resp, result.Name)
	}
	return
}

func GetPlanPagination(req view.GetPlanList) (resp view.PlanList, err error) {
	var page, pageSize, offset int
	var plans []db.Plan
	var count int
	if req.PageSize == 0 {
		pageSize = 1000
	} else {
		pageSize = int(req.PageSize)
	}
	if req.Page == 0 {
		page = 1
	} else {
		page = int(req.Page)
	}
	offset = (page - 1) * pageSize
	err = opt.DB.Offset(offset).Limit(pageSize).Find(&plans).Error
	if err != nil {
		return view.PlanList{}, err
	}
	err = opt.DB.Model(&db.Plan{}).Count(&count).Error
	if err != nil {
		return view.PlanList{}, err
	}
	for _, plan := range plans {
		resp.PlanList = append(resp.PlanList, view.PlanItem{
			Id:         plan.ID,
			Name:       plan.Name,
			Expression: plan.Expression,
			StartTime:  plan.StartTime,
			EndTime:    plan.EndTime,
			Period:     plan.Period,
		})
	}
	resp.CurrentPage = page
	resp.Total = count
	return
}

func CreatePlan(req view.CreatePlanReq) (err error) {
	//TODO add UserService
	var plan db.Plan
	err = opt.DB.Where("start_time=? and end_time=? and period=? and expression=?", req.StartTime, req.EndTime, req.Period, req.Expression).First(&plan).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return
		}
	} else {
		return errors.New("plan exist")
	}
	plan = db.Plan{
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Period:    req.Period,
		User:      "no_user",
	}
	err = opt.DB.Create(&plan).Error
	return
}

func ModifyPlan(req view.ModifyPlanReq) (err error) {
	var plan db.Plan
	err = opt.DB.Where("id=?", req.Id).First(&plan).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("plan not exist")
		}
	}
	err = opt.DB.Model(&db.Plan{}).Where("id=?", req.Id).Updates(&db.Plan{
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Period:     req.Period,
		Expression: req.Expression,
	}).Error
	return err
}

func DeletePlan(req view.DeleteRuleReq) (err error) {
	var plan db.Plan
	err = opt.DB.Model(&db.Plan{}).Where("id = ?", req.ID).Delete(&plan).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return err
	}
	return nil
}
