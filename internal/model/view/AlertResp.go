package view

type GetAlertResp struct {
	Pagination `json:"pagination"`
	Alerts     Alerts `json:"list"`
}

type Alerts []Alert

type Alert struct {
	ID       uint    `json:"id"`
	Labels   string  `json:"labels"`
	Value    float64 `json:"value"`
	Count    int     `json:"count"`
	Summary  string  `json:"summary"`
	Instance string  `json:"instance"`
	FiredAt  string  `json:"fired_at"`
}
