package view

type GetRule struct {
	Id uint `query:"id"`
}

type GetRulesReq struct {
	PaginationRequest
}

type CreateRuleReq struct {
	Expr        string `json:"expr" bind:"required"`
	Op          string `json:"op" bind:"required"`
	Value       string `json:"value" bind:"required"`
	For         string `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PlanName    string `json:"plan_name" bind:"required"`
	PromName    string `json:"prom_name" bind:"required"`
}

type ModifyRuleReq struct {
	ID          uint   `json:"id" bind:"required"`
	Expr        string `json:"expr" bind:"required"`
	Op          string `json:"op" bind:"required"`
	Value       string `json:"value" bind:"required"`
	For         string `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PlanName    string `json:"plan_name" bind:"required"`
	PromName    string `json:"prom_name" bind:"required"`
}

type DeleteRuleReq struct {
	ID uint `json:"id" binding:"required"`
}
