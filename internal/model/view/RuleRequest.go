package view

type GetRule struct {
	Id uint `query:"id"`
}

type GetRulesReq struct {
	Page     uint `json:"page"`
	PageSize uint `json:"page_size"`
}

type CreateRuleReq struct {
	Expr        string `json:"expr" bind:"required"`
	Op          string `json:"op" bind:"required"`
	Value       string `json:"value" bind:"required"`
	For         int    `json:"for" bind:"required"` //持续时间，单位秒
	Summary     string `json:"summary" bind:"required"`
	Description string `json:"description"`
	PlanName    string `json:"plan_name" bind:"required"`
	PromName    string `json:"prom_name" bind:"required"`
}

type ModifyRuleReq struct {
	ID          uint   `json:"id" bind:"required"`
	Expr        string `json:"expr" bind:"required"`
	Op          string `json:"op" bind:"required"`
	Value       string `json:"value" bind:"required"`
	For         int    `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PlanName    string `json:"plan_name" bind:"required"`
	PromName    string `json:"prom_name" bind:"required"`
}

type DeleteRuleReq struct {
	ID uint `json:"id" binding:"required"`
}
