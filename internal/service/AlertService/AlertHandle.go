/*
 * @Time : 2020/10/23 15:57
 * @Author : wangyl
 * @File : AlertHandle.go
 * @Software: GoLand
 */
package AlertService

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"time"
	"xing-doraemon/global"
	"xing-doraemon/internal/model/db"
	"xing-doraemon/internal/service/DingTalkService"
	"xing-doraemon/pkg/Utils"
)

var (
	defaultHandle AlertHandle
)

type AlertHandle interface {
	HandleAlert(task interface{})
}

func RegisterAlertHandle(alertHandle AlertHandle) {
	defaultHandle = alertHandle
}

func GetAlertHandle() AlertHandle {
	if defaultHandle == nil {
		return DefaultAlertHandle{}
	}
	return defaultHandle
}

type DefaultAlertHandle struct {
}

/*

 */
func (pThis DefaultAlertHandle) HandleAlert(task interface{}) {
	//TODO alert Handle
	data, ok := task.([]byte)
	if !ok {
		return
	}
	var alerts []PromAlertItem
	if err := json.Unmarshal(data, &alerts); err != nil {
		global.Log.Error("HandleAlert json.Unmarshal fail:" + err.Error())
		return
	}
	var dingTalkInfo DingTalkService.DingTalkInfo
	if len(alerts) > 0 {
		expr, err := GetExpression(uint(Utils.MustToInt(alerts[0].Annotations.RuleId)))
		if err != nil {
			global.Log.Error("HandleAlert GetExpression fail:" + err.Error())
			return
		}
		dingTalkInfo.Title = alerts[0].Annotations.Summary
		dingTalkInfo.Text = expr
	}
	for _, alert := range alerts {
		labes, send, err := HandleOneAlert(alert)
		if err != nil {
			continue
		}
		if send {
			dingTalkInfo.AlertList = append(dingTalkInfo.AlertList, DingTalkService.AlertItem{
				Labels: labes,
				Value:  alert.Value,
			})
		}
	}
	if err := DingTalkService.PushDingTalkInfo(dingTalkInfo); err != nil {
		global.Log.Error("HandleAlert PushDingTalkInfo fail:" + err.Error())
	}
}

func GetExpression(ruleID uint) (string, error) {
	var rule db.Rule
	err := opt.DB.Where("id=?", ruleID).Find(&rule).Error
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s %s", rule.Expr, rule.Op, rule.Value), nil
}

func HandleOneAlert(oneAlert PromAlertItem) (labels string, send bool, err error) {
	logger := global.Log
	var alert db.Alert
	labels, err = jsoniter.MarshalToString(oneAlert.Labels)
	if err != nil {
		logger.Error("HandleOneAlert json.Marshal Labels fail:" + err.Error())
		return
	}
	err = opt.DB.Where("rule_id=? and labels=?", oneAlert.Annotations.RuleId, labels).Find(&alert).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("HandleOneAlert Find db.Alert fail:" + err.Error())
		return
	}
	if alert.ConfirmedBy == "" {
		send = true
	}
	if alert.ConfirmedAt != nil {
		if time.Now().Sub(*alert.ConfirmedAt) > 2*time.Hour {
			send = true
		}
	}
	if alert.ID > 0 {
		count := alert.Count + 1
		err := opt.DB.Model(&db.Alert{}).Where("id=?", alert.ID).Updates(map[string]interface{}{
			"count":   count,
			"last_at": alert.LastAt,
		}).Error
		if err != nil {
			logger.Error("HandleOneAlert update count db.Alert fail:" + err.Error())
		}
	} else {
		alert = db.Alert{
			Labels:      labels,
			Value:       oneAlert.Value,
			Count:       1,
			Status:      oneAlert.State,
			Summary:     oneAlert.Annotations.Summary,
			Description: oneAlert.Annotations.Description,
			Instance:    oneAlert.Labels["instance"],
			FiredAt:     oneAlert.FiredAt,
			LastAt:      oneAlert.LastSentAt,
			RuleId:      uint(Utils.MustToInt(oneAlert.Annotations.RuleId)),
		}
		err = opt.DB.Create(&alert).Error
		if err != nil {
			logger.Error("HandleOneAlert update Create db.Alert fail:" + err.Error())
		}
	}
	return
}
