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
	StartTime uint `query:"start_time"`
	EndTime   uint `query:"end_time"`
	PaginationRequest
}

type CreatePlanReq struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Period    int    `json:"period" binding:"required"`
}

type ModifyPlanReq struct {
	Id        uint   `json:"id" binding:"required"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Period    int    `json:"period"`
}

type DeletePlanReq struct {
	Id uint `json:"id" binding:"required"`
}
