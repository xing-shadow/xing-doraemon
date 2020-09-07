/*
@Time : 2020/9/3 17:26
@Author : wangyl
@File : TimerService.go
@Software: GoLand
*/
package TimerService

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"

	"xing-doraemon/gobal"
	"xing-doraemon/interval/Invoker"
	"xing-doraemon/interval/model"
	"xing-doraemon/pkg/common"
)

type Record struct {
	Id              int64
	RuleId          int64
	Value           float64
	Count           int
	Summary         string
	Description     string
	Hostname        string
	ConfirmedBefore *time.Time
	FiredAt         *time.Time
	Labels          string
}

func (r Record) getLabelMap() map[string]string {
	labelMap := map[string]string{}
	if r.Labels != "" {
		for _, j := range strings.Split(r.Labels, "\v") {
			kv := strings.Split(j, "\a")
			labelMap[kv[0]] = kv[1]
		}
	}

	return labelMap
}

type RecoverRecord struct {
	Id       int64
	RuleId   int64
	Value    float64
	Count    int
	Summary  string
	Hostname string
}

func UpdateMaintainlist() {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			gobal.GetLogger().Errorf("Panic in UpdateMaintainlist:%v\n%s", err, buf)
		}
	}()
	delta, _ := time.ParseDuration("30s")
	datetime := time.Now().Add(delta)
	now := datetime.Format("15:04")
	maintainIds := []struct {
		Id int64
	}{}
	Invoker.GetDB().Table(model.Maintains{}.TableName()).Select("id").
		Where("valid>=? AND day_start<=? AND day_end>=? AND (flag=true AND (time_start<=? OR time_end>=?) OR flag=false AND time_start<=? AND time_end>=?) AND month&"+strconv.Itoa(int(math.Pow(2, float64(time.Now().Month()))))+">0", datetime.Format("2006-01-02 15:04:05"), datetime.Day(), datetime.Day(), now, now, now, now).Find(&maintainIds)
	//TODO
	m := map[string]bool{}
	for _, mid := range maintainIds {
		hosts := []struct {
			Hostname string
		}{}
		Invoker.GetDB().Table(model.Hosts{}.TableName()).
			Select("hostname").Where("mid = ?", mid.Id).Find(&hosts)
		for _, name := range hosts {
			m[name.Hostname] = true
		}
	}
	resp, err := common.HttpGet(gobal.GetAlterGatewayConfig().Send.BrokenUrl, nil, map[string]string{"Authorization": "Bearer 8gi6UvoPJgIRcunHBWDHel4fCLQVn9"})
	if err == nil {
		jsonDataFromHttp, _ := ioutil.ReadAll(resp.Body)
		brokenList := common.BrokenList{}
		json.Unmarshal(jsonDataFromHttp, &brokenList)
		for _, i := range brokenList.Hosts {
			m[i.Hostname] = true
		}
	}
	common.Rw.Lock()
	common.Maintain = m
	common.Rw.Unlock()
}

func Filter(alerts map[int64][]Record, maxCount map[int64]int) map[string][]common.Ready2Send {
	SendClass := map[string][]common.Ready2Send{
		common.AlertMethodSms:    {},
		common.AlertMethodLanxin: {},
		common.AlertMethodCall:   {},
		//"HOOK":   []common.Ready2Send{},
	}
	Cache := map[int64][]common.UserGroup{}
	NewRuleCount := map[[2]int64]int64{}
	for key := range alerts {
		var usergroupList []common.UserGroup
		var planId struct {
			PlanId  int64
			Summary string
		}
		AlertsMap := map[int][]common.SingleAlert{}
		Invoker.GetDB().Where(model.Rules{}.TableName()).Select("plan_id,summary").
			Where("id = ?", key).Find(&planId)
		if _, ok := Cache[planId.PlanId]; !ok {
			Invoker.GetDB().Table(model.Receivers{}.TableName()).Select("id,start_time,end_time,start,period,reverse_polish_notation,user,group,duty_group,method").
				Where("plan_id = ?", planId.PlanId).Find(&usergroupList)
			Cache[planId.PlanId] = usergroupList
		}
		for _, element := range Cache[planId.PlanId] {
			if element.IsValid() && element.IsOnDuty() {
				if maxCount[key] >= element.Start {
					k := [2]int64{key, int64(element.Start)}
					if _, ok := common.RuleCount[k]; !ok {
						NewRuleCount[k] = -1
					} else {
						NewRuleCount[k] = common.RuleCount[k]
					}
					NewRuleCount[k] += 1

					if NewRuleCount[k]%int64(element.Period) == 0 {
						// add alerts to AlertsMap
						if _, ok := AlertsMap[element.Start]; !ok {
							putToAlertMap(AlertsMap, element, alerts[key])
						}
						// forward alerts in AlertsMap to SendClass
						if len(AlertsMap[element.Start]) > 0 {
							var filteredAlerts []common.SingleAlert
							if element.ReversePolishNotation == "" {
								filteredAlerts = AlertsMap[element.Start]
							} else {
								for _, alert := range AlertsMap[element.Start] {
									if common.CalculateReversePolishNotation(alert.Labels, element.ReversePolishNotation) {
										filteredAlerts = append(filteredAlerts, alert)
									}
								}
							}
							putToSendClass(SendClass, key, element, filteredAlerts)
						}
					}
				}
			}
		}
	}
	common.RuleCount = NewRuleCount
	return SendClass
}

func putToSendClass(sendClass map[string][]common.Ready2Send, ruleId int64, ug common.UserGroup, alerts []common.SingleAlert) {
	if len(alerts) <= 0 {
		return
	}

	sendClass[ug.Method] = append(sendClass[ug.Method], common.Ready2Send{
		RuleId: ruleId,
		Start:  ug.Id,
		User: model.SendAlertsFor(Invoker.GetDB(), &common.ValidUserGroup{
			User:      ug.User,
			Group:     ug.Group,
			DutyGroup: ug.DutyGroup,
		}),
		Alerts: alerts,
	})
}

func putToAlertMap(alertMap map[int][]common.SingleAlert, ug common.UserGroup, alerts []Record) {

	alertMap[ug.Start] = []common.SingleAlert{}

	for _, alert := range alerts {
		if alert.Count >= ug.Start {
			if _, ok := common.Maintain[alert.Hostname]; !ok {
				alertMap[ug.Start] = append(alertMap[ug.Start], common.SingleAlert{
					Id:       alert.Id,
					Count:    alert.Count,
					Value:    alert.Value,
					Summary:  alert.Summary,
					Hostname: alert.Hostname,
					Labels:   alert.getLabelMap(),
				})
			}
		}
	}
}

func InitTimerService() {
	go func() {
		for {
			current := time.Now()
			time.Sleep(time.Duration(90-current.Second()) * time.Second)
			UpdateMaintainlist()
		}
	}()
	go func() {
		for {
			current := time.Now()
			time.Sleep(time.Duration(60-current.Second()) * time.Second)
			now := time.Now().Format("2006-01-02 15:04:05")
			go func() {
				defer func() {
					if e := recover(); e != nil {
						buf := make([]byte, 16384)
						buf = buf[:runtime.Stack(buf, false)]
						gobal.GetLogger().Errorf("Panic in timer:%v\n%s", e, buf)
					}
				}()
				var info []Record
				Invoker.GetDB().Model(&model.Alerts{}).Update(map[string]string{
					"state": "2",
				}).Where("status = ? AND confirmed_before < ?", 1, now).Find(&info)
				tx := Invoker.GetDB().Begin()
				tx.Exec("UPDATE alert SET count=count+1 WHERE status!=0")
				tx.Table(model.Alerts{}.TableName()).Select("id,rule_id,value,count,summary,description,hostname,confirmed_before,fired_at,labels").
					Where("status = ? ", 2).Find(&info)
				//filter alters
				aggregation := map[int64][]Record{}
				maxCount := map[int64]int{}
				for _, i := range info {
					aggregation[i.RuleId] = append(aggregation[i.RuleId], i)
					if _, ok := maxCount[i.RuleId]; !ok {
						maxCount[i.RuleId] = i.Count
					} else {
						if i.Count > maxCount[i.RuleId] {
							maxCount[i.RuleId] = i.Count
						}
					}
				}
				common.Rw.RLock()
				ready2send := Filter(aggregation, maxCount)
				common.Rw.RUnlock()
				tx.Commit()
				gobal.GetLogger().Infof("Alerts to send:%v", ready2send)
				Sender(ready2send, now)
				common.Lock.Lock()
				recover2send := common.Recover2Send
				common.Recover2Send = map[string]map[[2]int64]*common.Ready2Send{
					common.AlertMethodLanxin: {},
					//"HOOK":   map[[2]int64]*common.Ready2Send{},
				}
				common.Lock.Unlock()
				gobal.GetLogger().Info("Recoveries to send:%v", recover2send)
				RecoverSender(recover2send, now)
			}()
		}
	}()
}
