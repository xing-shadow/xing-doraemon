package AlertService

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/prometheus/rules"
	"runtime"
	"time"
	"xing-doraemon/global"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/model/view"
	"xing-doraemon/internal/service/AlertService/Sender"
	"xing-doraemon/pkg/xtime"
)

func HandleAlerts(alerts []view.Alert) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			global.Log.Error(fmt.Sprintf("Panic in AlertsHandler:%v\n%s", e, buf))
		}
	}()
	timeZero := xtime.ZeroTime()
	for _, item := range alerts {
		var alert db.Alert
		err := opt.DB.Where("rule_id =? AND labels=? AND fired_at=?", item.Annotations.RuleId, item.Labels, &item.FiredAt).First(&alert).Error
		if err == nil {
			// alert has been triggered
			if alert.State != int(rules.StateInactive) {
				if item.State == int(rules.StateInactive) {
					if err := recoverAlert(alert.ID, item); err != nil {
						global.Log.Error("recoverAlert fail:" + err.Error())
						continue
					}
				} else {
					opt.DB.Model(&db.Alert{}).Where("rule_id =? AND labels=? AND fired_at=?", item.Annotations.RuleId, item.Labels, &item.FiredAt).Updates(map[string]interface{}{
						"summary":     item.Annotations.Summary,
						"description": item.Annotations.Description,
						"value":       item.Value,
					})
				}
			} else {
				continue
			}
		} else if gorm.IsRecordNotFoundError(err) {
			alert = db.Alert{
				Labels:          item.Labels.String(),
				Value:           item.Value,
				State:           item.State,
				Summary:         item.Annotations.Summary,
				Description:     item.Annotations.Description,
				Hostname:        item.Labels.Get("instance"),
				FiredAt:         &item.FiredAt,
				ConfirmedAt:     &timeZero,
				ConfirmedBefore: &timeZero,
				ResolvedAt:      &timeZero,
				RuleId:          item.Annotations.RuleId,
			}
		}
	}
}

func recoverAlert(id uint, alert view.Alert) error {
	if id <= 0 {
		return errors.New("Invalid Param ")
	}
	err := opt.DB.Model(&db.Alert{}).Where("id = ?", id).Updates(map[string]interface{}{
		"state":       rules.StateInactive,
		"resolved_at": &alert.ResolvedAt,
		"summary":     alert.Annotations.Summary,
		"description": alert.Annotations.Description,
		"value":       alert.Value,
	}).Error
	if err != nil {
		return err
	}
	var rule db.Rule
	err = opt.DB.Preload("Plan").Select("id").Where("id=?", alert.Annotations.RuleId).Find(&rule).Error
	if err != nil {
		return err
	}
	recoveryMutex.Lock()
	if _, ok := recoverySend[rule.Plan.Method]; !ok {
		recoverySend[rule.Plan.Method] = &Sender.ReadyToSend{
			RulId: rule.ID,
			Url:   rule.Plan.Url,
		}
	}
	recoverySend[rule.Plan.Method].Alerts = append(recoverySend[rule.Plan.Method].Alerts, Sender.SendAlert{
		Id:       id,
		Value:    alert.Value,
		Summary:  alert.Annotations.Summary,
		Hostname: alert.Labels.Get("instance"),
		Labels:   alert.Labels.String(),
	})
	recoveryMutex.Unlock()
	return nil
}

func filterAlert(alerts map[uint][]db.Alert) (result map[int]*Sender.ReadyToSend) {
	result = make(map[int]*Sender.ReadyToSend)
	now := time.Now()
	for key := range alerts {
		var plan db.Plan
		var rule db.Rule
		err := opt.DB.Select("plan_id").Where("id = ?", key).Find(&rule).Error
		if err != nil {
			global.Log.Debug("filterAlert get rule fail:" + err.Error())
			continue
		}
		err = opt.DB.Where("id =?", rule.PlanID).First(&plan).Error
		if err != nil {
			global.Log.Debug("filterAlert get plan fail:" + err.Error())
			continue
		}
		if plan.IsSend(now) {
			result[plan.Method] = &Sender.ReadyToSend{
				RulId: rule.ID,
				Url:   plan.Url,
			}
			for _, item := range alerts[key] {
				if item.ConfirmedBefore != nil {
					if now.Before(*item.ConfirmedBefore) {
						continue
					}
				}
				result[plan.Method].Alerts = append(result[plan.Method].Alerts, Sender.SendAlert{
					Id:       item.ID,
					Value:    item.Value,
					Summary:  item.Summary,
					Hostname: item.Hostname,
					Labels:   item.Labels,
				})
			}
		}
	}
	return
}
