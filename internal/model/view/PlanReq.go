package view

type GetPlan struct {
	Id uint `query:"id"`
}

type GetPlanList struct {
	Page     uint `query:"page"`
	PageSize uint `query:"page_size"`
}

type CreatePlanReq struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Method    int    `json:"method"`
	Url       string `json:"url"`
}

type ModifyPlanReq struct {
	Id        uint   `json:"id" binding:"required"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Method    int    `json:"method"`
	Url       string `json:"url"`
}

type DeletePlanReq struct {
	Id uint `json:"id" binding:"required"`
}
