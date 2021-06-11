package view

type GetPlan struct {
	Id uint `query:"id"`
}

type GetPlanList struct {
	StartTime uint `query:"start_time"`
	EndTime   uint `query:"end_time"`
	Page      uint `json:"page"`
	PageSize  uint `json:"page_size"`
}

type CreatePlanReq struct {
	Name       string `json:"name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Expression string `json:"expression"`
	Period     int    `json:"period" binding:"required"`
}

type ModifyPlanReq struct {
	Id         uint   `json:"id" binding:"required"`
	Name       string `json:"name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Expression string `json:"expression"`
	Period     int    `json:"period"`
}

type DeletePlanReq struct {
	Id uint `json:"id" binding:"required"`
}
