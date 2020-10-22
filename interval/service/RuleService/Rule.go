package RuleService

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"xing-doraemon/global"
	"xing-doraemon/interval/model/db"
	"xing-doraemon/interval/model/view"
)

func GetRule(req view.GetRule) (resp view.RuleItem, err error) {
	var rule db.Rule
	err = opt.DB.Where("id=?", req.Id).First(&rule).Error
	if err != nil {
		return
	}
	resp = view.RuleItem{
		Id:          rule.ID,
		Expr:        rule.Expr,
		Op:          rule.Op,
		Value:       rule.Value,
		For:         rule.For,
		Summary:     rule.Summary,
		Description: rule.Description,
		PlanID:      rule.PlanID,
		PromID:      rule.PromID,
	}
	return
}

func GetPaginationRule(req view.GetRulesReq) (resp view.RuleList, err error) {
	var page, offset, pageSize int
	var rules []db.Rule
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
	err = opt.DB.Find(&rules).Offset(offset).Limit(pageSize).Error
	if err != nil {
		global.GetLogger().Error("get rules from mysql fail:", err)
		return
	}
	err = opt.DB.Model(&db.Rule{}).Count(&count).Error
	if err != nil {
		global.GetLogger().Error("get rules counts from mysql fail:", err)
		return
	}
	resp.Total = count
	resp.CurrentPage = page
	for _, rule := range rules {
		var rule = view.RuleItem{
			Id:          rule.ID,
			Expr:        rule.Expr,
			Op:          rule.Op,
			Value:       rule.Value,
			For:         rule.For,
			Summary:     rule.Summary,
			Description: rule.Description,

			PlanID: rule.PlanID,
			PromID: rule.PromID,
		}
		resp.Rules = append(resp.Rules, rule)
	}
	return
}

func GetAllRule() (resp view.RuleList, err error) {
	var rules []db.Rule
	var count int
	err = opt.DB.Find(&rules).Error
	if err != nil {
		global.GetLogger().Error("get rules from mysql fail:", err)
		return
	}
	err = opt.DB.Model(&db.Rule{}).Count(&count).Error
	if err != nil {
		global.GetLogger().Error("get rules counts from mysql fail:", err)
		return
	}
	resp.Total = count
	for _, rule := range rules {
		var rule = view.RuleItem{
			Id:          rule.ID,
			Expr:        rule.Expr,
			Op:          rule.Op,
			Value:       rule.Value,
			For:         rule.For,
			Summary:     rule.Summary,
			Description: rule.Description,

			PlanID: rule.PlanID,
			PromID: rule.PromID,
		}
		resp.Rules = append(resp.Rules, rule)
	}
	return
}

func CreateRule(req view.CreateRuleReq) (err error) {
	var prom db.Prom
	var plan db.Plan
	var rule db.Rule
	err = opt.DB.Where("id = ?", req.Prom).Find(&prom).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fmt.Errorf("not found prom record")
		}
		return err
	}
	err = opt.DB.Where("id = ?", req.PlanID).Find(&plan).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fmt.Errorf("not found plan record")
		}
		return err
	}
	err = opt.DB.Where("expr=? and op=? and value=?", req.Expr, req.Op, req.Value).First(&rule).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return
		}
	} else {
		return fmt.Errorf("record exist")
	}
	rule = db.Rule{
		PlanID:      plan.ID,
		PromID:      prom.ID,
		Expr:        req.Expr,
		Op:          req.Op,
		Value:       req.Value,
		For:         req.For,
		Summary:     req.Summary,
		Description: req.Description,
	}
	err = opt.DB.Create(&rule).Error
	return
}

func ModifyRule(req view.ModifyRuleReq) (err error) {
	var rule db.Rule

	err = opt.DB.Where("id=?", req.ID).First(&rule).Error
	if err != nil {
		return
	}
	err = opt.DB.Model(&rule).Where("id=?", req.ID).Updates(&db.Rule{
		Expr:        req.Expr,
		Op:          req.Op,
		Value:       req.Value,
		For:         req.For,
		Summary:     req.Summary,
		Description: req.Summary,
	}).Error
	return
}

func DeleteRule(req view.DeleteRuleReq) (err error) {
	var rule db.Rule
	err = opt.DB.Where("id=?", req.ID).Delete(&rule).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return err
	}
	return nil
}
