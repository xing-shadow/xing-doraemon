package view

type RuleList struct {
	Pagination `json:"pagination"`
	Rules      []RuleItem `json:"list"`
}

type RuleItem struct {
	Id          uint   `json:"id"`
	Expr        string `json:"expr"`
	Op          string `json:"op"`
	Value       string `json:"value"`
	For         int    `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`

	PlanName string `json:"plan_name"`
	PromName string `json:"prom_name"`
}
