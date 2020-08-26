/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : plan_receiver.go
@Software: GoLand
*/
package model

type Receivers struct {
	Id                    int64  `orm:"auto" json:"id,omitempty"`
	Plan                  *Plans `orm:"index;rel(fk)" json:"plan_id"`
	StartTime             string `orm:"size(31)" json:"start_time"`
	EndTime               string `orm:"size(31)" json:"end_time"`
	Start                 int    `json:"start"`
	Period                int    `json:"period"`
	Expression            string `orm:"size(1023)" json:"expression"`
	ReversePolishNotation string `orm:"size(1023)" json:"reverse_polish_notation"`
	User                  string `orm:"size(1023)" json:"user"`
	Group                 string `orm:"size(1023)" json:"group"`
	DutyGroup             string `orm:"size(255)" json:"duty_group"`
	Method                string `orm:"size(255)" json:"method"`
}

type Rec struct {
	Id         int64  `json:"id,omitempty"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Start      int    `json:"start"`
	Period     int    `json:"period"`
	Expression string `json:"expression"`
	User       string `json:"user"`
	Group      string `json:"group"`
	DutyGroup  string `json:"duty_group"`
	Method     string `json:"method"`
}

func (*Receivers) TableName() string {
	return "plan_receiver"
}
