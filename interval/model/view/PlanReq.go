/*
 * @Time : 2020/10/22 14:46
 * @Author : wangyl
 * @File : PlanReq.go
 * @Software: GoLand
 */
package view

type GetPlan struct {
	Id uint `query:"id"`
}

type GetPlanList struct {
	PaginationRequest
}

type CreatePlanReq struct {
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Period     int    `json:"period" binding:"required"`
	Expression string `json:"expression"`
}

type ModifyPlanReq struct {
	Id         uint   `json:"id" binding:"required"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Period     int    `json:"period"`
	Expression string `json:"expression"`
}

type DeletePlanReq struct {
	Id uint `json:"id" binding:"required"`
}
