package view

type RuleList struct {
	PaginationResp
	Rules []RuleItem `json:"rules"`
}

type RuleItem struct {
	Id          uint   `json:"id"`
	Expr        string `json:"expr"`
	Op          string `json:"op"`
	Value       string `json:"value"`
	For         string `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`

	PlanID uint `json:"plan_id"`
	PromID uint `json:"prom_id"`
}