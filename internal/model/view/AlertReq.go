package view

type GetAlertsReq struct {
	Page     uint `json:"page"`
	PageSize uint `json:"page_size"`
}

type ConfirmAlertsReq struct {
	AlertList []uint `json:"alert_list"`
}
