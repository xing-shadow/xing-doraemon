package view

type PlanList struct {
	Pagination `json:"pagination"`
	PlanList   []PlanItem `json:"list"`
}

type PlanItem struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Period     int    `json:"period"`
	Expression string `json:"expression"`
}
