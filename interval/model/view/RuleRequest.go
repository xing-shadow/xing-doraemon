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
	PlanID      uint   `json:"plan_id" bind:"required"`
	Prom        uint   `json:"prom" bind:"required"`
}

type ModifyRuleReq struct {
	ID          uint   `json:"id" bind:"required"`
	Expr        string `json:"expr" bind:"required"`
	Op          string `json:"op" bind:"required"`
	Value       string `json:"value" bind:"required"`
	For         string `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type DeleteRuleReq struct {
	ID uint `json:"id" binding:"required"`
}
