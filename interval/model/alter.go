/*
@Time : 2020/8/25 10:37
@Author : wangyl
@File : alter.go
@Software: GoLand
*/
package model

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"xing-doraemon/gobal"
	"xing-doraemon/pkg/common"
)

type Alerts struct {
	Id              int64      `gorm:"column:id;auto" json:"id,omitempty"`
	Rule            *Rules     `gorm:"rel:fk" json:"rule_id"`
	Labels          string     `gorm:"column:labels;size:4095" json:"labels"`
	Value           float64    `gorm:"column:value" json:"value"`
	Count           int        `json:"count"`
	Status          int8       `gorm:"index" json:"status"`
	Summary         string     `gorm:"column:summary;size:1023" json:"summary"`
	Description     string     `gorm:"column:description;size:1023" json:"description"`
	Hostname        string     `gorm:"column:hostname;size:255" json:"hostname"`
	ConfirmedBy     string     `gorm:"column:confirmed_by;size:1023" json:"confirmed_by"`
	FiredAt         *time.Time `gorm:"type:datetime" json:"fired_at"`
	ConfirmedAt     *time.Time `gorm:"null" json:"confirmed_at"`
	ConfirmedBefore *time.Time `gorm:"null" json:"confirmed_before"`
	ResolvedAt      *time.Time `gorm:"null" json:"resolved_at"`
}

type OneAlert struct {
	ID              int64      `json:"id"`
	RuleID          int64      `json:"rule_id"`
	Value           float64    `json:"value"`
	Status          int8       `json:"status"`
	Count           int        `json:"count"`
	Summary         string     `json:"summary"`
	Description     string     `json:"description"`
	ConfirmedBy     string     `json:"confirmed_by"`
	FiredAt         *time.Time `json:"fired_at"`
	ConfirmedAt     *time.Time `json:"confirmed_at"`
	ConfirmedBefore *time.Time `json:"confirmed_before"`
	ResolvedAt      *time.Time `json:"resolved_at"`
}

type ShowAlerts struct {
	Alerts []common.AlertForShow `json:"alerts"`
	Total  int64                 `json:"total"`
}

func (Alerts) TableName() string {
	return "alert"
}

type record struct {
	Id              int64
	RuleId          int64
	Labels          string
	Value           float64
	Count           int
	Status          int8
	Summary         string
	Description     string
	ConfirmedBy     string
	FiredAt         *time.Time
	ConfirmedAt     *time.Time
	ConfirmedBefore *time.Time
	ResolvedAt      *time.Time
}

func (r record) toOneAlert() OneAlert {
	return OneAlert{
		ID:              r.Id,
		RuleID:          r.RuleId,
		Value:           r.Value,
		Status:          r.Status,
		Count:           r.Count,
		Summary:         r.Summary,
		Description:     r.Description,
		ConfirmedBy:     r.ConfirmedBy,
		FiredAt:         r.FiredAt,
		ConfirmedAt:     r.ConfirmedAt,
		ConfirmedBefore: r.ConfirmedBefore,
		ResolvedAt:      r.ResolvedAt,
	}
}

func (r record) getLabelMap() map[string]string {
	label := map[string]string{}
	if r.Labels != "" {
		for _, e := range strings.Split(r.Labels, "\v") {
			kv := strings.Split(e, "\a")
			label[kv[0]] = kv[1]
		}
	}
	return label
}

func (r record) toAlertForShow() common.AlertForShow {

	return common.AlertForShow{
		Id:              r.Id,
		RuleId:          r.RuleId,
		Labels:          r.getLabelMap(),
		Value:           r.Value,
		Count:           r.Count,
		Status:          r.Status,
		Summary:         r.Summary,
		Description:     r.Description,
		ConfirmedBy:     r.ConfirmedBy,
		FiredAt:         r.FiredAt,
		ConfirmedAt:     r.ConfirmedAt,
		ConfirmedBefore: r.ConfirmedBefore,
		ResolvedAt:      r.ResolvedAt,
	}
}

func (u *Alerts) GetAlerts(db *gorm.DB, pageNo int64, pageSize int64, timeStart string, timeEnd string, status string, summary string) ShowAlerts {
	var showAlerts ShowAlerts
	showAlerts.Alerts = []common.AlertForShow{}
	var records []record
	if summary != "" {
		if status != "" {
			if timeStart != "" {
				if timeEnd != "" {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND fired_at<=? AND status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				} else {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("fired_at<=? AND status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeEnd, status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			} else {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("status=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", status, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			}
		} else {
			if timeStart != "" {
				if timeEnd != "" {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND fired_at<=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				} else {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeStart, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("fired_at<=? AND summary LIKE ? ORDER BY id DESC LIMIT ?,?", timeEnd, "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			} else {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("summary LIKE ? ORDER BY id DESC LIMIT ?,?", "%"+summary+"%", (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			}
		}
	} else {
		if status != "" {
			if timeStart != "" {
				if timeEnd != "" {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND fired_at<=? AND status=? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, status, (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				} else {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND status=? ORDER BY id DESC LIMIT ?,?", timeStart, status, (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("fired_at<=? AND status=? ORDER BY id DESC LIMIT ?,?", timeEnd, status, (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			} else {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("status=? ORDER BY id DESC LIMIT ?,?", status, (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			}
		} else {
			if timeStart != "" {
				if timeEnd != "" {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? AND fired_at<=? ORDER BY id DESC LIMIT ?,?", timeStart, timeEnd, (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				} else {
					db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
						Where("fired_at>=? ORDER BY id DESC LIMIT ?,?", timeStart, (pageNo-1)*pageSize, pageSize).
						Find(&records).Count(&showAlerts.Total)
				}
			} else if timeEnd != "" {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Where("fired_at<=? ORDER BY id DESC LIMIT ?,?", timeEnd, (pageNo-1)*pageSize, pageSize).
					Find(&records).Count(&showAlerts.Total)
			} else {
				db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
					Order("id DESC").Offset((pageNo - 1) * pageSize).Limit(pageSize).
					Find(&records).Count(&showAlerts.Total)
			}
		}
	}
	for _, i := range records {
		showAlerts.Alerts = append(showAlerts.Alerts, i.toAlertForShow())
	}
	return showAlerts
}

func (u *Alerts) ShowAlerts(db *gorm.DB, ruleId string, start string, pageNo int64, pageSize int64) ShowAlerts {
	var showAlerts ShowAlerts
	showAlerts.Alerts = []common.AlertForShow{}
	var records []record
	strategy := struct {
		ReversePolishNotation string
		Start                 int
	}{}
	if start != "" {
		db.Table(Receivers{}.TableName()).Select("start,reverse_polish_notation").Where("id=?", start).Find(&strategy)
	}
	db.Table(u.TableName()).Select("id,rule_id,labels,value,count,status,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
		Where("count>=? AND rule_id=? AND status!=0 ORDER BY status DESC,id DESC", strategy.Start, ruleId).Find(&records)
	for _, i := range records {
		label := i.getLabelMap()
		if strategy.ReversePolishNotation != "" {
			if common.CalculateReversePolishNotation(label, strategy.ReversePolishNotation) {
				showAlerts.Alerts = append(showAlerts.Alerts, i.toAlertForShow())
			}
		} else {
			showAlerts.Alerts = append(showAlerts.Alerts, i.toAlertForShow())
		}
	}
	showAlerts.Total = int64(len(showAlerts.Alerts))
	if showAlerts.Total == 0 {
		return showAlerts
	} else if showAlerts.Total < pageNo*pageSize {
		showAlerts.Alerts = showAlerts.Alerts[(pageNo-1)*pageSize:]
		return showAlerts
	} else {
		showAlerts.Alerts = showAlerts.Alerts[(pageNo-1)*pageSize : pageNo*pageSize]
		return showAlerts
	}
}

func (u *Alerts) ClassifyAlerts(db *gorm.DB) map[string]map[string][]OneAlert {
	var records []record
	db.Table(u.TableName()).Select("id,rule_id,labels,value,status,count,summary,description,confirmed_by,fired_at,confirmed_at,confirmed_before,resolved_at").
		Where("status=? AND count!=?", 2, -1).Find(&records)
	res := map[string]map[string][]OneAlert{}
	for _, i := range records {
		if i.Labels != "" {
			for _, j := range strings.Split(i.Labels, "\v") {
				kv := strings.Split(j, "\a")
				if _, ok := res[kv[0]]; ok {
					res[kv[0]][kv[1]] = append(res[kv[0]][kv[1]], i.toOneAlert())
				} else {
					res[kv[0]] = map[string][]OneAlert{}
					res[kv[0]][kv[1]] = append(res[kv[0]][kv[1]], i.toOneAlert())
				}
			}
		} else {
			if _, ok := res["no label"]; ok {
				res["no label"]["no label"] = append(res["no label"]["no label"], i.toOneAlert())
			} else {
				res["no label"] = map[string][]OneAlert{}
				res["no label"]["no label"] = append(res["no label"]["no label"], i.toOneAlert())
			}
		}
	}
	return res
}

func (u *Alerts) ConfirmAll(db *gorm.DB, confirmList *common.Confirm) error {
	now := time.Now()
	var err error
	for _, id := range confirmList.Ids {
		var rs struct {
			Status uint8
		}
		tx := db.Begin()
		err = tx.Table(u.TableName()).Select("status,rule_id").Where("id=? LOCK IN SHARE MODE", id).Find(&rs).Error
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "database query error")
		} else {
			const AlertStatusOn = 2
			if rs.Status == AlertStatusOn {
				tx.Model(&Alerts{}).Where("id = ? ", id).Updates(map[string]interface{}{
					"status":           1,
					"confirmed_at":     now.Format("2006-01-02 15:04:05"),
					"confirmed_by":     confirmList.User,
					"confirmed_before": now.Add(time.Duration(confirmList.Duration) * time.Minute).Format("2006-01-02 15:04:05"),
				})
			}
		}
		tx.Commit()
	}
	return errors.Wrap(err, "database update error")
}

func (u *Alerts) AlertsHandler(db *gorm.DB, alert *common.Alerts) {
	logger := gobal.GetLogger()
	Cache := map[int64][]common.UserGroup{}
	todayZero, _ := time.ParseInLocation("2006-01-02", "2019-01-01 15:22:22", time.Local)
	for _, elemt := range *alert {
		var queryres []struct {
			Id     int64
			Status uint8
		}
		a := &alertForQuery{Alert: &elemt}
		a.setFields()
		err := db.Table(u.TableName()).Select("id,status").
			Where("rule_id =? AND labels=? AND fired_at=?", a.ruleId, a.label, a.firedAt).Find(&queryres).Error
		if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			if len(queryres) > 0 {
				// alert has been triggered by post requests before
				if queryres[0].Status != 0 {
					const AlertStatusOff = 0
					if elemt.State == AlertStatusOff {
						// handle the recover message
						recoverAlert(db, *a, Cache)
					} else {
						db.Model(&Alerts{}).Where("rule_id =? AND labels=? AND fired_at=?", a.ruleId, a.label, a.firedAt).
							Update(map[string]interface{}{
								"summary":     elemt.Annotations.Summary,
								"description": elemt.Annotations.Description,
								"value":       elemt.Value,
							})
					}
				} else {
					continue
				}
			} else {
				// insert an new alert
				var alert Alerts
				alert.Id = 0 //reset the "Id" to 0,which is very important:after a record is inserted,the value of "Id" will not be 0,but the auto primary key of the record
				alert.Rule = &Rules{Id: a.ruleId}
				alert.Labels = a.label
				alert.FiredAt = &a.firedAt
				alert.Description = elemt.Annotations.Description
				alert.Summary = elemt.Annotations.Summary
				alert.Count = -1
				alert.Value = elemt.Value
				alert.Status = int8(elemt.State)
				alert.Hostname = a.hostname
				alert.ConfirmedAt = &todayZero
				alert.ConfirmedBefore = &todayZero
				alert.ResolvedAt = &todayZero
				err := db.Create(&alert).Error
				if err != nil {
					logger.Error("Insert alter failed:%s", err)
				}
			}
		}
	}
}

type alertForQuery struct {
	*common.Alert
	label    string
	hostname string
	ruleId   int64
	firedAt  time.Time
}

/*
 set value for fields in alertForQuery
*/
func (a *alertForQuery) setFields() {
	var orderKey []string
	var labels []string

	// set ruleId
	a.ruleId, _ = strconv.ParseInt(a.Annotations.RuleId, 10, 64)
	for key := range a.Labels {
		orderKey = append(orderKey, key)
	}
	sort.Strings(orderKey)
	for _, i := range orderKey {
		labels = append(labels, i+"\a"+a.Labels[i])
	}
	// set label
	a.label = strings.Join(labels, "\v")
	// set firedAt
	a.firedAt = a.FiredAt.Truncate(time.Second)
	// set hostname
	a.setHostname()
}

/*
 set hostname by instance label on data
*/
func (a *alertForQuery) setHostname() {
	h := ""
	if _, ok := a.Labels["instance"]; ok {
		h = a.Labels["instance"]
		boundary := strings.LastIndex(h, ":")
		if boundary != -1 {
			h = h[:boundary]
		}
	}
	a.hostname = h
}

/*
 process for receiving an recovery alert:
*/
func recoverAlert(db *gorm.DB, a alertForQuery, cache map[int64][]common.UserGroup) {
	recoverInfo := struct {
		Id       int64
		Count    int
		Hostname string
	}{}

	tx := db.Begin()
	err := db.Table(Alerts{}.TableName()).Select("id,count,hostname").
		Where("rule_id =? AND labels=? AND fired_at=? FOR UPDATE", a.ruleId, a.label, a.firedAt).Find(&recoverInfo).Error
	if err == nil {
		if recoverInfo.Id != 0 {
			updadeErr := tx.Model(&Alerts{}).Where("id=?", recoverInfo.Id).Updates(map[string]interface{}{
				"status":      a.State,
				"summary":     a.Annotations.Summary,
				"description": a.Annotations.Description,
				"value":       a.Value,
				"resolved_at": a.ResolvedAt,
			}).Error
			if updadeErr == nil {
				common.Rw.RLock()
				if _, ok := common.Maintain[a.hostname]; !ok {
					var userGroupList []common.UserGroup
					var planId struct {
						PlanId  int64
						Summary string
					}
					tx.Table(Rules{}.TableName()).Where("id=?", a.ruleId).Find(&planId)
					if _, ok := cache[planId.PlanId]; !ok {
						db.Table(Receivers{}.TableName()).Select("id,start_time,end_time,start,period,reverse_polish_notation,user,group,duty_group,method").
							Where("plan_id=? AND (method='LANXIN' OR method LIKE 'HOOK %')", planId.PlanId).Find(&userGroupList)
						cache[planId.PlanId] = userGroupList
					}
					for _, element := range cache[planId.PlanId] {

						if !(element.IsValid() && element.IsOnDuty()) {
							continue
						}

						if !(recoverInfo.Count >= element.Start) {
							continue
						}

						if ok := shouldSend(recoverInfo.Id, a.ruleId, recoverInfo.Count, element); !ok {
							continue
						}

						if element.ReversePolishNotation != "" && !common.CalculateReversePolishNotation(a.Labels, element.ReversePolishNotation) {
							continue
						}

						// merge users
						users := SendAlertsFor(db, &common.ValidUserGroup{
							User:      element.User,
							Group:     element.Group,
							DutyGroup: element.DutyGroup,
						})

						// update Recover2Send, other goroutines in timer.go will handle it
						common.UpdateRecovery2Send(element, *a.Alert, users, recoverInfo.Id, recoverInfo.Count, recoverInfo.Hostname)
					}
				}
				common.Rw.RUnlock()
				tx.Commit()
			} else {
				tx.Rollback()
			}
		}
		tx.Commit()
	} else {
		tx.Rollback()
		db.Exec("UPDATE alert SET status=?,summary=?,description=?,value=?,resolved_at=? WHERE id=?", a.State, a.Annotations.Summary, a.Annotations.Description, a.Value, a.ResolvedAt, recoverInfo.Id)
	}
}

/*
 whether recovery should be send
*/
func shouldSend(alertId, ruleId int64, alertCount int, ug common.UserGroup) (sendFlag bool) {
	//sendFlag := false
	if alertCount-ug.Start >= ug.Period {
		sendFlag = true
	} else {
		if _, ok := common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]; ok {
			gobal.GetLogger().Debugf("id:%d,rulecount:%d,count:%d,start:%d,period:%d", alertId, common.RuleCount[[2]int64{ruleId, int64(ug.Start)}], alertCount, ug.Start, ug.Period)
			if common.RuleCount[[2]int64{ruleId, int64(ug.Start)}] >= int64(alertCount-ug.Start) {
				gobal.GetLogger().Debugf("[%s] id:%d %d,%s", alertId, (common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]-int64(alertCount)+int64(ug.Start))%int64(ug.Period), common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]-((common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]-int64(alertCount)+int64(ug.Start))/int64(ug.Period))*int64(ug.Period) >= int64(ug.Period))
				if (common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]-int64(alertCount)+int64(ug.Start))%int64(ug.Period) == 0 || common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]-((common.RuleCount[[2]int64{ruleId, int64(ug.Start)}]-int64(alertCount)+int64(ug.Start))/int64(ug.Period))*int64(ug.Period) >= int64(ug.Period) {
					sendFlag = true
				}
			}
		}
	}

	return
}
