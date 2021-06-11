package AlertService

import (
	"time"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/model/view"
)

func GetAlertList(req view.GetAlertsReq) (resp view.GetAlertResp, err error) {
	var alerts []db.Alert
	var page, pageSize, offset int
	var count int
	if req.Page <= 0 {
		page = 1
	} else {
		page = int(req.Page)
	}
	if req.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = int(req.PageSize)
	}
	offset = (page - 1) * pageSize
	err = opt.DB.Where("status=? and confirmed_at is null", 2).
		Or("status=? and confirmed_at is not null and ? > confirmed_at", 2, time.Now().Add(-2*time.Hour)).
		Offset(offset).Limit(pageSize).Find(&alerts).Error
	if err != nil {
		return view.GetAlertResp{}, err
	}
	err = opt.DB.Table(db.Alert{}.TableName()).Where("status=? and confirmed_at is null", 2).Count(&count).Error
	if err != nil {
		return view.GetAlertResp{}, err
	}
	resp.Total = count
	resp.CurrentPage = int(page)
	resp.PageSize = pageSize
	for _, alert := range alerts {
		resp.Alerts = append(resp.Alerts, view.Alert{
			ID:       alert.ID,
			Labels:   alert.Labels,
			Value:    alert.Value,
			Count:    alert.Count,
			Summary:  alert.Summary,
			Instance: alert.Instance,
			FiredAt:  alert.FiredAt.Format(time.RFC3339),
		})
	}
	return
}

func ConfirmAlertList(userName string, req view.ConfirmAlertsReq) (err error) {
	err = opt.DB.Model(&db.Alert{}).Where("id in (?)", req.AlertList).Updates(map[string]interface{}{
		"confirmed_by": userName,
		"confirmed_at": time.Now(),
	}).Error
	return
}
