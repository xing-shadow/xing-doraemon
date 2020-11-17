/*
 * @Time : 2020/11/5 16:05
 * @Author : wangyl
 * @File : Struct.go
 * @Software: GoLand
 */
package AlertService

import (
	"time"
)

/*
[
    {
        "active_at": "2020-11-05T15:40:40.489299029+08:00",
        "annotations": {
            "description": "携程数量过大",
            "prom_id": "2",
            "rule_id": "1",
            "summary": "携程指标"
        },
        "fired_at": "2020-11-05T15:41:40.489299029+08:00",
        "labels": {
            "instance": "localhost:9090",
            "job": "prometheus"
        },
        "last_sent_at": "2020-11-05T15:41:40.489299029+08:00",
        "resolved_at": "0001-01-01T00:00:00Z",
        "state": 2,
        "valid_until": "2020-11-05T15:44:40.489299029+08:00",
        "value": 31
    }
]
*/
type PromAlertItem struct {
	ActiveAt    *time.Time        `json:"active_at"`    //触发时间
	FiredAt     *time.Time        `json:"fired_at"`     //第一次告警时间
	State       int8              `json:"state"`        //状态 0:正常 1:挂起 2:告警
	LastSentAt  string            `json:"last_sent_at"` //上一次告警时间
	Value       float64           `json:"value"`        //上一次触法告警值
	ValidUntil  string            `json:"valid_until"`  //有效期时间节点
	Annotations Annotations       `json:"annotations"`  //元数据
	Labels      map[string]string `json:"labels"`       //标签列表
}

type Annotations struct {
	Description string `json:"description"` //详细描述
	PromId      string `json:"prom_id"`     //prom记录id
	RuleId      string `json:"rule_id"`     //rule记录id
	Summary     string `json:"summary"`     //简介
}
