package AlertService

import (
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/rules"
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
	if req.State != "" {
		if req.StartTime != "" {
			if req.EndTime != "" {
				err = opt.DB.Table(db.Alert{}.TableName()).Where("state=? AND fired_at < ? AND fired_at > ?", req.State, req.EndTime, req.StartTime).Count(&count).Offset(offset).Limit(pageSize).Find(&alerts).Error
				if err != nil {
					return
				}
			} else {
				err = opt.DB.Table(db.Alert{}.TableName()).Where("state=? AND fired_at < ?", req.State, req.EndTime).Count(&count).Offset(offset).Limit(pageSize).Find(&alerts).Error
				if err != nil {
					return
				}
			}
		} else {
			err = opt.DB.Table(db.Alert{}.TableName()).Where("status=?", req.State).Count(&count).Offset(offset).Limit(pageSize).Find(&alerts).Error
			if err != nil {
				return
			}
		}
	} else {
		err = opt.DB.Table(db.Alert{}.TableName()).Count(&count).Offset(offset).Limit(pageSize).Find(&alerts).Error
		if err != nil {
			return
		}
	}
	resp.Total = count
	resp.CurrentPage = int(page)
	resp.PageSize = pageSize
	for _, alert := range alerts {
		resp.Alerts = append(resp.Alerts, view.AlertItem{
			ID:              alert.ID,
			RuleId:          alert.RuleId,
			Labels:          alert.Labels,
			Value:           alert.Value,
			Status:          alert.State,
			Summary:         alert.Summary,
			Instance:        alert.Hostname,
			Description:     alert.Description,
			ConfirmedBy:     alert.ConfirmedBy,
			FiredAt:         alert.FiredAt,
			ConfirmedAt:     alert.ConfirmedAt,
			ConfirmedBefore: alert.ConfirmedBefore,
			ResolvedAt:      alert.ResolvedAt,
		})
	}
	return
}

func ConfirmAlertList(userName string, req view.ConfirmAlertsReq) (err error) {
	now := time.Now()
	confirmedBefore := time.Now().Add(time.Duration(req.Duration) * time.Minute)
	for i := 0; i < len(req.AlertList); i++ {
		var alert db.Alert
		tx := opt.DB.Begin()
		err = tx.Select("rule_id,status").Where("id=?", req.AlertList[i]).First(&alert).Error
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "database query error")
		} else {
			if alert.State == int(rules.StateFiring) {
				err = tx.Model(&db.Alert{}).Where("id=?", req.AlertList[i]).Updates(map[string]interface{}{
					"status":           1,
					"confirmed_by":     userName,
					"confirmed_at":     &now,
					"confirmed_before": &confirmedBefore,
				}).Error
				if err != nil {
					tx.Rollback()
					return errors.Wrap(err, "database update error")
				}
			}
		}
		tx.Commit()
	}
	return
}
