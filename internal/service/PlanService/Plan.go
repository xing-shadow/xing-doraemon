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
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}
	if plan.ID <= 0 {
		err = errors.New("改告警计划不存在")
		return
	}
	resp = view.PlanItem{
		Id:        plan.ID,
		Name:      plan.Name,
		StartTime: plan.StartTime,
		EndTime:   plan.EndTime,
		Method:    plan.Method,
		Url:       plan.Url,
	}
	return
}

func GetPlanAllName() (resp []string, err error) {
	var results []db.Plan
	err = opt.DB.Select("name").Find(&results).Error
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
	err = opt.DB.Select("id, name, start_time, end_time, method, url").Offset(offset).Limit(pageSize).Find(&plans).Error
	if err != nil {
		return view.PlanList{}, err
	}
	err = opt.DB.Model(&db.Plan{}).Count(&count).Error
	if err != nil {
		return view.PlanList{}, err
	}
	for _, plan := range plans {
		resp.PlanList = append(resp.PlanList, view.PlanItem{
			Id:        plan.ID,
			Name:      plan.Name,
			StartTime: plan.StartTime,
			EndTime:   plan.EndTime,
			Method:    plan.Method,
			Url:       plan.Url,
		})
	}
	resp.CurrentPage = page
	resp.Total = count
	resp.PageSize = pageSize
	return
}

func CreatePlan(req view.CreatePlanReq) (err error) {
	var plan db.Plan
	plan = db.Plan{
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Method:    req.Method,
		Url:       req.Url,
	}
	err = opt.DB.Save(&plan).Error
	return
}

func ModifyPlan(req view.ModifyPlanReq) (err error) {
	var plan db.Plan
	err = opt.DB.Where("id=?", req.Id).First(&plan).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("改告警计划不存在")
			return
		}
		return
	}
	plan.Name = req.Name
	plan.Url = req.Name
	plan.StartTime = req.StartTime
	plan.EndTime = req.EndTime
	plan.Method = req.Method
	plan.Url = req.Url
	return err
}

func DeletePlan(req view.DeleteRuleReq) (err error) {
	var plan db.Plan
	err = opt.DB.Where("id = ?", req.ID).Delete(&plan).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("改告警计划不存在")
			return
		}
		return err
	}
	return nil
}
