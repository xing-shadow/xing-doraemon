/*
 * @Time : 2020/10/22 14:47
 * @Author : wangyl
 * @File : PlanResp.go
 * @Software: GoLand
 */
package view

type PlanList struct {
	PaginationResp
	PlanList []PlanItem `json:"plan_list"`
}

type PlanItem struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Period    int    `json:"period"`
}
