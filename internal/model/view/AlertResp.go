/*
 * @Time : 2020/11/13 16:30
 * @Author : wangyl
 * @File : AlertResp.go
 * @Software: GoLand
 */
package view

type GetAlertResp struct {
	PaginationResp
	Alerts Alerts `json:"alerts"`
}

type Alerts []Alert

type Alert struct {
	ID       uint    `json:"id"`
	Labels   string  `json:"labels"`
	Value    float64 `json:"value"`
	Count    int     `json:"count"`
	Summary  string  `json:"summary"`
	Instance string  `json:"instance"`
	FiredAt  string  `json:"fired_at"`
}
